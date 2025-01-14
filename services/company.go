package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/pkg/mail-service"
	"github.com/kalleriakronos24/khaimal-group/templates/email"
	"github.com/kalleriakronos24/khaimal-group/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type CheckExistingCompanyStruct struct {
	*masterModels.Company
}

func (module *module) InsertCompany(c *gin.Context, p *dto.InsertCompany, pAgent *dto.DriverSignUpValidator, pDrivers []*dto.DriverSignUpValidator, pTransports []*dto.InsertCarManagement) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	// ==== INSERT COMPANY ====
	var company *masterModels.Company
	if company, err = module.db.companyModel.InsertCompany(masterModels.Company{
		Name:           p.Name,
		PhoneNumber:    p.PhoneNumber,
		Email:          p.Email,
		CompanyName:    p.CompanyName,
		CompanyAddress: p.CompanyAddress,
		CompanyCountry: p.CompanyCountry,
		CompanyNumber:  p.CompanyNumber,
		Status:         "pending-approval",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}
	// ==== END INSERT COMPANY ====

	// ==== INSERT MAIN AGENT USER DRIVER ====
	pAgentDto := &dto.DriverSignup{
		Name:     pAgent.Name,
		Email:    pAgent.Email,
		Phone:    pAgent.Phone,
		Password: "12345678",
	}

	var hashedPasswordAgent []byte
	if hashedPasswordAgent, err = bcrypt.GenerateFromPassword([]byte(pAgentDto.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("server error. please try again later")
	}

	var credAgent *masterModels.Credentials
	if credAgent, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    pAgentDto.Email,
		Password: string(hashedPasswordAgent),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = module.db.userDriverModel.InsertDriver(masterModels.Driver{
		Name:          pAgentDto.Name,
		CredentialsID: credAgent.ID,
		DriverType:    "internal-agent",
		Phone:         pAgentDto.Phone,
		CompanyID:     company.ID,
		Status:        "inactive",
		Prob:          80,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	var driverAgent masterModels.Credentials
	if txError := tx.Model(&driverAgent).
		Where("id = ?", credAgent.ID).
		Preload(clause.Associations).
		First(&driverAgent); txError.Error != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = module.db.balanceDriver.InsertBalanceDriver(models.BalanceDriver{
		Amount:   30,
		DriverID: driverAgent.CredentialDriver.ID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	// ===== INSERT CAR TO ALL THE DRIVER REGISTERED =====
	if len(pTransports) > 0 {
		for idxTransport := 0; idxTransport < len(pTransports); idxTransport++ {
			randomFileName, _ := utils.GenerateNumber(30)
			uniqueFileName := fmt.Sprintf("%v%v", randomFileName, ".jpg")
			carModelid := pTransports[idxTransport].CarModelID
			// if pTransports[idxTransport].FrontCarPhoto != nil {
			// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "front-car-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
			// 		return errors.New("failed to get create car management")
			// 	}
			// }

			// if pTransports[idxTransport].LicensePhoto != nil {
			// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "license-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
			// 		return errors.New("failed to get create car management")
			// 	}
			// }

			// if pTransports[idxTransport].RoadTaxPhoto != nil {
			// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "road-tax-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
			// 		return errors.New("failed to get create car management")
			// 	}
			// }

			// if pTransports[idxTransport].VEPPhoto != nil {
			// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "vep-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
			// 		return errors.New("failed to get create car management")
			// 	}
			// }

			CarManagement := masterModels.CarManagement{
				Name:              pTransports[idxTransport].Name,
				PlateNumber:       pTransports[idxTransport].PlateNumber,
				CompanyID:         company.ID,
				LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
				CarManagementType: pTransports[idxTransport].CarManagementType,
				DriverID:          driverAgent.CredentialDriver.ID,
				CarModelID:        carModelid,
				FileName:          uniqueFileName,
				Status:            "inactive",
				FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
				RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
				VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			}

			// if pTransports[idxTransport].LicensePhoto == nil {
			// 	CarManagement = masterModels.CarManagement{
			// 		CompanyID:         company.ID,
			// 		Name:              pTransports[idxTransport].Name,
			// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
			// 		CarManagementType: pTransports[idxTransport].CarManagementType,
			// 		DriverID:          driverAgent.CredentialDriver.ID,
			// 		CarModelID:        carModelid,
			// 		FileName:          uniqueFileName,
			// 		Status:            "inactive",
			// 		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 	}
			// }

			// if pTransports[idxTransport].FrontCarPhoto == nil {
			// 	CarManagement = masterModels.CarManagement{
			// 		CompanyID:         company.ID,
			// 		Name:              pTransports[idxTransport].Name,
			// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
			// 		LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		CarManagementType: pTransports[idxTransport].CarManagementType,
			// 		DriverID:          driverAgent.CredentialDriver.ID,
			// 		CarModelID:        carModelid,
			// 		FileName:          uniqueFileName,
			// 		Status:            "inactive",
			// 		RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 	}
			// }

			// if pTransports[idxTransport].VEPPhoto == nil {
			// 	CarManagement = masterModels.CarManagement{
			// 		CompanyID:         company.ID,
			// 		Name:              pTransports[idxTransport].Name,
			// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
			// 		LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		CarManagementType: pTransports[idxTransport].CarManagementType,
			// 		DriverID:          driverAgent.CredentialDriver.ID,
			// 		CarModelID:        carModelid,
			// 		FileName:          uniqueFileName,
			// 		Status:            "inactive",
			// 		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 	}
			// }

			// if pTransports[idxTransport].RoadTaxPhoto == nil {
			// 	CarManagement = masterModels.CarManagement{
			// 		CompanyID:         company.ID,
			// 		Name:              pTransports[idxTransport].Name,
			// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
			// 		LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		CarManagementType: pTransports[idxTransport].CarManagementType,
			// 		DriverID:          driverAgent.CredentialDriver.ID,
			// 		CarModelID:        carModelid,
			// 		FileName:          uniqueFileName,
			// 		Status:            "inactive",
			// 		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 		VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			// 	}
			// }

			if CarManagementErr := tx.Create(&CarManagement); CarManagementErr.Error != nil {
				tx.Rollback()
				return errors.New("failed to upload car management")
			}
		}
	}
	// ==== END INSERT CAR MANAGEMENT ====
	// ==== END INSERT MAIN AGENT ====

	// ==== INSERT DRIVER AGENT =====
	if len(pDrivers) > 0 {
		for i := 0; i < len(pDrivers); i++ {
			fDriverLicensePhoto, _ := c.FormFile("drivers.licensePhoto")

			pDriver := &dto.DriverSignup{
				Name:         pDrivers[i].Name,
				Email:        pDrivers[i].Email,
				Phone:        pDrivers[i].Phone,
				DriverType:   "internal",
				LicensePhoto: fDriverLicensePhoto,
				Password:     "12345678",
				Prob:         0,
			}

			if fDriverLicensePhoto != nil {
				pDriver = &dto.DriverSignup{
					Name:         pDrivers[i].Name,
					Email:        pDrivers[i].Email,
					Phone:        pDrivers[i].Phone,
					DriverType:   "internal",
					LicensePhoto: fDriverLicensePhoto,
					Password:     "12345678",
					Prob:         0,
				}
			}

			if fDriverLicensePhoto == nil {
				pDriver = &dto.DriverSignup{
					Name:       pDrivers[i].Name,
					Email:      pDrivers[i].Email,
					Phone:      pDrivers[i].Phone,
					DriverType: "internal",
					Password:   "12345678",
					Prob:       0,
				}
			}
			var hashedPassword []byte
			if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(pDriver.Password), bcrypt.DefaultCost); err != nil {
				return errors.New("server error. please try again later")
			}

			var cred *masterModels.Credentials
			if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
				Email:    pDriver.Email,
				Password: string(hashedPassword),
			}, tx); err != nil {
				tx.Rollback()
				return errors.New("failed to register. try again")
			}

			if err = module.db.userDriverModel.InsertDriver(masterModels.Driver{
				Name:          pDriver.Name,
				CredentialsID: cred.ID,
				DriverType:    "internal",
				Phone:         pDriver.Phone,
				CompanyID:     company.ID,
				Status:        "inactive",
				Prob:          0,
			}, tx); err != nil {
				tx.Rollback()
				return errors.New("failed to register. try again")
			}

			var driver masterModels.Credentials
			if txError := tx.Model(&driver).
				Where("id = ?", cred.ID).
				Preload(clause.Associations).
				First(&driver); txError.Error != nil {
				tx.Rollback()
				return errors.New("failed to register. try again")
			}
			if err = module.db.balanceDriver.InsertBalanceDriver(models.BalanceDriver{
				Amount:   0,
				DriverID: driver.CredentialDriver.ID,
			}, tx); err != nil {
				tx.Rollback()
				return errors.New("failed to register. try again")
			}

			// ===== INSERT CAR TO ALL THE DRIVER REGISTERED =====
			if len(pTransports) > 0 {
				for idxTransport := 0; idxTransport < len(pTransports); idxTransport++ {
					randomFileName, _ := utils.GenerateNumber(30)
					uniqueFileName := fmt.Sprintf("%v%v", randomFileName, ".jpg")
					carModelid := pTransports[idxTransport].CarModelID
					// if pTransports[idxTransport].FrontCarPhoto != nil {
					// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "front-car-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
					// 		return errors.New("failed to get create car management")
					// 	}
					// }

					// if pTransports[idxTransport].LicensePhoto != nil {
					// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "license-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
					// 		return errors.New("failed to get create car management")
					// 	}
					// }

					// if pTransports[idxTransport].RoadTaxPhoto != nil {
					// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "road-tax-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
					// 		return errors.New("failed to get create car management")
					// 	}
					// }

					// if pTransports[idxTransport].VEPPhoto != nil {
					// 	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "vep-photo", pTransports[idxTransport].FrontCarPhoto, uniqueFileName); saveFileErr != nil {
					// 		return errors.New("failed to get create car management")
					// 	}
					// }
					CarManagement := masterModels.CarManagement{
						CompanyID:         company.ID,
						Name:              pTransports[idxTransport].Name,
						PlateNumber:       pTransports[idxTransport].PlateNumber,
						LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
						CarManagementType: pTransports[idxTransport].CarManagementType,
						DriverID:          driver.CredentialDriver.ID,
						CarModelID:        carModelid,
						FileName:          uniqueFileName,
						Status:            "inactive",
						FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
						RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
						VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					}

					// if pTransports[idxTransport].LicensePhoto == nil {
					// 	CarManagement = masterModels.CarManagement{
					// 		CompanyID:         company.ID,
					// 		Name:              pTransports[idxTransport].Name,
					// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
					// 		CarManagementType: pTransports[idxTransport].CarManagementType,
					// 		DriverID:          driver.CredentialDriver.ID,
					// 		CarModelID:        carModelid,
					// 		FileName:          uniqueFileName,
					// 		Status:            "inactive",
					// 		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 	}
					// }

					// if pTransports[idxTransport].FrontCarPhoto == nil {
					// 	CarManagement = masterModels.CarManagement{
					// 		CompanyID:         company.ID,
					// 		Name:              pTransports[idxTransport].Name,
					// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
					// 		LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		CarManagementType: pTransports[idxTransport].CarManagementType,
					// 		DriverID:          driver.CredentialDriver.ID,
					// 		CarModelID:        carModelid,
					// 		FileName:          uniqueFileName,
					// 		Status:            "inactive",
					// 		RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 	}
					// }

					// if pTransports[idxTransport].VEPPhoto == nil {
					// 	CarManagement = masterModels.CarManagement{
					// 		CompanyID:         company.ID,
					// 		Name:              pTransports[idxTransport].Name,
					// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
					// 		LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		CarManagementType: pTransports[idxTransport].CarManagementType,
					// 		DriverID:          driver.CredentialDriver.ID,
					// 		CarModelID:        carModelid,
					// 		FileName:          uniqueFileName,
					// 		Status:            "inactive",
					// 		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 	}
					// }

					// if pTransports[idxTransport].RoadTaxPhoto == nil {
					// 	CarManagement = masterModels.CarManagement{
					// 		CompanyID:         company.ID,
					// 		Name:              pTransports[idxTransport].Name,
					// 		PlateNumber:       pTransports[idxTransport].PlateNumber,
					// 		LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		CarManagementType: pTransports[idxTransport].CarManagementType,
					// 		DriverID:          driver.CredentialDriver.ID,
					// 		CarModelID:        carModelid,
					// 		FileName:          uniqueFileName,
					// 		Status:            "inactive",
					// 		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 		VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
					// 	}
					// }

					if CarManagementErr := tx.Create(&CarManagement); CarManagementErr.Error != nil {
						tx.Rollback()
						return errors.New("failed to upload car management")
					}
				}
			}
			// ==== END INSERT CAR MANAGEMENT ====
		}
	}
	// ==== END INSERT DRIVER AGENT ====

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  driverAgent.Email,
		Subject: "Company Registration Submitted",
		Body: `
		Please wait our approval

		it might takes 1-2 business day, and we will email you back once we approved.


		thanks for be part of WadahGo
		`,
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}

func (module *module) RetrieveCompanyByUserID(userId uuid.UUID) (m masterModels.Company, err error) {
	if m, err = module.db.companyModel.GetOneCompanyByUserID(userId); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) ApproveCompany(companyId uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	var company masterModels.Company
	if company, err = module.db.companyModel.GetOneCompanyByID(companyId); err != nil {
		return errors.New(err.Error())
	}

	if err := module.db.companyModel.UpdateCompany(company.ID, masterModels.Company{
		Status: "active",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.carManagementModel.UpdateCarManagementByCompanyId(company.ID, masterModels.CarManagement{
		Status: "active",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	if err := module.db.driverModel.UpdateDriverByCompanyId(company.ID, masterModels.Driver{
		Status: "active",
	}, tx); err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	var mainAgentDriver masterModels.Driver
	if mainAgentDriver, err = module.db.driverModel.GetOneMainAgentByCompanyId(companyId); err != nil {
		return errors.New(err.Error())
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  mainAgentDriver.Credentials.Email,
		Subject: "Company Registration Success",
		Body: `
		Your Company is approved!

		Please note that, all default password for the drivers and the Company PIC is: 12345678
		`,
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return err
}

func (module *module) RetrieveAllCompany(id uuid.UUID) (m []masterModels.Company, err error) {
	if m, err = module.db.companyModel.GetAllCompany(id); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RetrieveAllCompanyPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.companyModel.GetAllCompanyPaginated(c, id); err != nil {
		return pagination, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RetrieveCompany(id uuid.UUID) (m masterModels.Company, err error) {
	if m, err = module.db.companyModel.GetOneCompanyByID(id); err != nil {
		return m, fmt.Errorf("%s", err.Error())
	}
	return
}

func (module *module) RegisterNewDriverInternalAgent(driverId uuid.UUID, credentials *dto.DriverSignup) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {
		return errors.New("server error. please try again later")
	}

	log.Printf("manager id >> %v", driverId)

	var driverManager masterModels.Driver
	if txError := tx.Model(&driverManager).
		Where("credentials_id = ?", driverId).
		Preload(clause.Associations).
		First(&driverManager); txError.Error != nil {
		tx.Rollback()
		return errors.New("failed to register new driver. try again")
	}

	var cred *masterModels.Credentials

	if cred, err = module.db.credentialModel.InsertCredentials(masterModels.Credentials{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	if err = module.db.userDriverModel.InsertDriver(masterModels.Driver{
		Name:          credentials.Name,
		CredentialsID: cred.ID,
		DriverType:    "internal",
		Phone:         credentials.Phone,
		Prob:          0,
		CompanyID:     driverManager.CompanyID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register new driver. try again")
	}

	var driver masterModels.Credentials
	if txError := tx.Model(&driver).
		Where("id = ?", cred.ID).
		Preload(clause.Associations).
		First(&driver); txError.Error != nil {
		tx.Rollback()
		return errors.New("failed to register new driver. try again")
	}

	if err = module.db.balanceDriver.InsertBalanceDriver(models.BalanceDriver{
		Amount:   0,
		DriverID: driver.CredentialDriver.ID,
	}, tx); err != nil {
		tx.Rollback()
		return errors.New("failed to register. try again")
	}

	var carManagement []masterModels.CarManagement
	if txError := tx.Model(&carManagement).
		Where("driver_id = ?", driverManager.ID).
		Preload(clause.Associations).
		Find(&carManagement); txError.Error != nil {
		tx.Rollback()
		return errors.New("failed to register new driver. try again")
	}

	randomFileName, _ := utils.GenerateNumber(30)
	uniqueFileName := fmt.Sprintf("%v%v", randomFileName, ".jpg")

	for i := range carManagement {
		p := carManagement[i]
		CarManagement := masterModels.CarManagement{
			Name:              p.Name,
			PlateNumber:       p.PlateNumber,
			LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			CarManagementType: p.CarManagementType,
			DriverID:          driver.CredentialDriver.ID,
			CarModelID:        p.CarModelID,
			FileName:          uniqueFileName,
			CompanyID:         p.CompanyID,
			FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
			VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
		}

		if CarManagementErr := tx.Create(&CarManagement); CarManagementErr.Error != nil {
			tx.Rollback()
			return errors.New("failed to upload car management")
		}
	}

	if err = mail.SendMailV3(&mail.TSendMail{
		From:    "WadahGo <notification@wadahgo.com>",
		MailTo:  cred.Email,
		Subject: "Registration Success",
		Body:    email.ETRegisterSuccess(credentials.Name),
	}); err != nil {
		return errors.New("server error. please try again later")
	}
	tx.Commit()
	return
}
