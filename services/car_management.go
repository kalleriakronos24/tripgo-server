package services

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

type CheckExistingCarManagementStruct struct {
	*master.CarManagement
}

func (module *module) InsertCarManagement(c *gin.Context, p *dto.InsertCarManagement) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	randomFileName, _ := utils.GenerateNumber(30)
	fileExt := filepath.Ext(p.FrontCarPhoto.Filename)
	uniqueFileName := fmt.Sprintf("%v%v", randomFileName, fileExt)

	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "front-car-photo", p.FrontCarPhoto, uniqueFileName); saveFileErr != nil {
		return errors.New("failed to get create car management")
	}

	CarManagement := master.CarManagement{
		Name:              p.Name,
		PlateNumber:       p.PlateNumber,
		LicensePhoto:      p.LicensePhoto,
		CarManagementType: p.CarManagementType,
		DriverID:          p.DriverID,
		CarModelID:        p.CarModelID,
		FileName:          uniqueFileName,
		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
	}
	if CarManagementErr := tx.Create(&CarManagement); CarManagementErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to upload car management")
	}
	tx.Commit()
	return
}

func (module *module) RetrieveCarManagementrByUserID(userId uuid.UUID) (m []master.CarManagement, err error) {
	if m, err = module.db.carManagementModel.GetAllCarManagementByDriverID(userId); err != nil {
		return nil, errors.New("failed to get car management list")
	}
	return m, nil
}

func (module *module) UpdateCarManagement(c *gin.Context, p *dto.UpdateCarManagement, id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	randomFileName, _ := utils.GenerateNumber(30)
	fileExt := filepath.Ext(p.FrontCarPhoto.Filename)
	uniqueFileName := fmt.Sprintf("%v%v", randomFileName, fileExt)

	var carManagement master.CarManagement
	if carManagement, err = module.db.carManagementModel.GetOneByID(id); err != nil {
		return errors.New("failed to get car management detail")
	}

	if _, saveFileErr := utils.UpdateFileFromDockerVolume(c, "car-management", "front-car-photo", p.FrontCarPhoto, carManagement.FileName, uniqueFileName); saveFileErr != nil {
		return errors.New("failed to update car management")
	}

	CarManagement := master.CarManagement{
		Name:          p.Name,
		PlateNumber:   p.PlateNumber,
		LicensePhoto:  p.LicensePhoto,
		CarModelID:    p.CarModelID,
		FileName:      uniqueFileName,
		FrontCarPhoto: fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
	}
	if CarManagementErr := tx.Model(&master.CarManagement{}).Where("id", id).Updates(&CarManagement); CarManagementErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to update car management")
	}
	tx.Commit()
	return
}

func (module *module) DeleteCarManagement(c *gin.Context, id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	var carManagement master.CarManagement
	if carManagement, err = module.db.carManagementModel.GetOneByID(id); err != nil {
		return errors.New("failed to get car management detail")
	}
	if _, saveFileErr := utils.RemoveFileFromDockerVolume(c, "car-management", "front-car-photo", carManagement.FileName); saveFileErr != nil {
		return errors.New("failed to delete car management")
	}
	if CarManagementErr := tx.Model(&master.CarManagement{}).Delete(&master.CarManagement{}, id); CarManagementErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to delete car management")
	}
	tx.Commit()
	return
}

func (module *module) SetCarManagementInactive(c *gin.Context, id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	CarManagement := master.CarManagement{
		Status: "inactive",
	}
	if CarManagementErr := tx.Model(&master.CarManagement{}).Where("id", id).Updates(&CarManagement); CarManagementErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to set car management to inactive")
	}
	tx.Commit()
	return
}

func (module *module) SetCarManagementActive(c *gin.Context, id uuid.UUID) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	CarManagement := master.CarManagement{
		Status: "active",
	}

	if CarManagementErr := tx.Model(&master.CarManagement{}).Where("id", id).Updates(&CarManagement); CarManagementErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to set car management to active")
	}
	tx.Commit()
	return
}
