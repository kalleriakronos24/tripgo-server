package services

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	masterModels "github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

type CheckExistingCompanyStruct struct {
	*masterModels.Company
}

func (module *module) InsertCompany(c *gin.Context, p *dto.InsertCarManagement) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	randomFileName, _ := utils.GenerateNumber(30)
	// fileExt := filepath.Ext(p.FrontCarPhoto.Filename)
	uniqueFileName := fmt.Sprintf("%v%v", randomFileName, ".jpg")

	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "front-car-photo", p.FrontCarPhoto, uniqueFileName); saveFileErr != nil {
		return errors.New("failed to get create car management")
	}

	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "license-photo", p.FrontCarPhoto, uniqueFileName); saveFileErr != nil {
		return errors.New("failed to get create car management")
	}

	if p.RoadTaxPhoto != nil {
		if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "road-tax-photo", p.FrontCarPhoto, uniqueFileName); saveFileErr != nil {
			return errors.New("failed to get create car management")
		}
	}

	if p.VEPPhoto != nil {
		if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "vep-photo", p.FrontCarPhoto, uniqueFileName); saveFileErr != nil {
			return errors.New("failed to get create car management")
		}
	}

	CarManagement := master.CarManagement{
		Name:              p.Name,
		PlateNumber:       p.PlateNumber,
		LicensePhoto:      fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
		CarManagementType: p.CarManagementType,
		DriverID:          p.DriverID,
		CarModelID:        p.CarModelID,
		FileName:          uniqueFileName,
		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
		RoadTaxPhoto:      fmt.Sprintf("%s/%s/car-management/road-tax-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
		VEPPhoto:          fmt.Sprintf("%s/%s/car-management/vep-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
	}

	if CarManagementErr := tx.Create(&CarManagement); CarManagementErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to upload car management")
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
