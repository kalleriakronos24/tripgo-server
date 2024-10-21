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
		FrontCarPhoto:     fmt.Sprintf("%s/%s/%s/car-management/front-car-photo/%s", config.AppConfig.APPUrl, config.AppConfig.APPUrlStaticFileGroupRoute, config.AppConfig.AppUrlStaticFileMainRoute, p.FrontCarPhoto.Filename),
	}

	if CarManagementErr := tx.Create(&CarManagement); CarManagementErr.Error != nil {
		tx.Rollback()
		return CarManagementErr.Error
	}

	tx.Commit()
	return
}

// func (module *module) SaveManagementDocument(id uuid.UUID) (output GenerateDocumentOutput, err error) {

// 	subTotalProductFormatted := formatCurrency.FormatMoney(subTotalProduct)
// 	ppnProduct = formatCurrency.FormatMoney(subTotalProduct * 0.11)
// 	grandTotalProduct = formatCurrency.FormatMoney(subTotalProduct + (subTotalProduct * 0.11))

// 	date := utils.ConvertEnToIDDateTime(CarManagement.Date)

// 	var outputPath string
// 	var saveFileErr error
// 	if outputPath, saveFileErr = utils.SaveFileToDockerVolume(nil, "", "po-out", nil); saveFileErr != nil {
// 		return output, errors.New(saveFileErr.Error())
// 	}

// 	data := GenerateDocumentOutput{
// 		OutputPath: outputPath,
// 		FileName:   CarManagement.Number,
// 	}

// 	return data, nil
// }
