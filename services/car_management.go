package services

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
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

	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "car-management", "front-car-photo", p.FrontCarPhoto); saveFileErr != nil {
		return errors.New(saveFileErr.Error())
	}
	CarManagement := master.CarManagement{
		Name:              p.Name,
		PlateNumber:       p.PlateNumber,
		LicensePhoto:      p.LicensePhoto,
		CarManagementType: p.CarManagementType,
		DriverID:          p.DriverID,
		CarModelID:        p.CarModelID,
		FrontCarPhoto:     fmt.Sprintf("%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.AppUrlStaticFileMainRoute, p.FrontCarPhoto.Filename),
	}

	if CarManagementErr := tx.Create(&CarManagement); CarManagementErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to upload car management")
	}
	tx.Commit()
	return
}
