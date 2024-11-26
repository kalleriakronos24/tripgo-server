package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kalleriakronos24/khaimal-group/config"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/models"
	"github.com/kalleriakronos24/khaimal-group/models/master"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

type CheckExistingDriverTopupStruct struct {
	*models.DriverTopup
}

func (module *module) RetrieveAllDriverTopup() (m []models.DriverTopup, err error) {
	if m, err = module.db.driverTopupModel.GetAllTopupRequest(); err != nil {
		return m, errors.New("failed to get driver topup requests")
	}
	return
}

func (module *module) RetrieveAllDriverTopupByDriver(id uuid.UUID) (m []models.DriverTransactionHistory, err error) {
	if m, err = module.db.driverTransactionHistoryModel.GetAllHistoryByDriverID(id); err != nil {
		return m, errors.New("failed to get transaction history")
	}
	return
}

func (module *module) InsertDriverTopup(c *gin.Context, p *dto.InsertDriverTopup) (err error) {
	tx := database.GetDatabaseConnection().Begin()
	randomFileName, _ := utils.GenerateNumber(30)
	uniqueFileName := fmt.Sprintf("%v%v", randomFileName, ".jpg")

	if _, saveFileErr := utils.SaveFileToDockerVolume(c, "driver-topup", "uploaded-receipt", p.ReceiptImage, uniqueFileName); saveFileErr != nil {
		return errors.New("failed to upload balance topup")
	}

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	month := int(currentMonth)
	randomUid, _ := utils.GenerateNumber(10)

	// config.AppConfig.APPUrl
	DriverTopup := models.DriverTopup{
		Uid:          fmt.Sprintf("DRV/TP/%v%v/%v", utils.IntegerToRoman(currentYear), utils.IntegerToRoman(month), randomUid),
		Amount:       float64(p.Amount),
		ReceiptImage: fmt.Sprintf("%s/%s/driver-topup/uploaded-receipt/%s", "http://localhost:3009/api/v1", config.AppConfig.AppUrlStaticFileMainRoute, uniqueFileName),
		DriverID:     p.DriverID,
	}

	if DriverTopupErr := tx.Create(&DriverTopup); DriverTopupErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to upload balance topup")
	}

	DriverTransactionHistory := models.DriverTransactionHistory{
		Remark:        "waiting approval",
		Status:        "waiting approval",
		DriverID:      p.DriverID,
		DriverTopupID: DriverTopup.ID,
	}

	if DriverTransactionHistoryErr := tx.Create(&DriverTransactionHistory); DriverTransactionHistoryErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to upload balance topup")
	}

	tx.Commit()
	return
}

func (module *module) ApproveDriverTopup(c *gin.Context, p *dto.UpdateDriverTopupHistory) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	DriverTransactionHistory := models.DriverTransactionHistory{
		Remark:        "approved",
		Status:        "approved",
		DriverID:      p.DriverID,
		DriverTopupID: p.DriverTopupID,
	}

	if DriverTransactionHistoryErr := tx.Model(&models.DriverTransactionHistory{}).Where("driver_topup_id", p.DriverTopupID).Updates(&DriverTransactionHistory); DriverTransactionHistoryErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to approve driver topup")
	}

	// update the driver balance
	var balanceDriver models.BalanceDriver
	if balanceDriver, err = module.db.balanceDriver.GetOneByID(p.DriverID); err != nil {
		return errors.New("failed to get wallet information")
	}

	BalanceDriverPayload := models.BalanceDriver{
		Amount: balanceDriver.Amount + float64(p.Amount),
	}

	if BalanceDriverErr := tx.Model(&models.BalanceDriver{}).Where("id", balanceDriver.ID).Updates(&BalanceDriverPayload); BalanceDriverErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to approve driver topup")
	}

	tx.Commit()
	return err
}

func (module *module) RejectDriverTopup(c *gin.Context, p *dto.UpdateDriverTopupHistory) (err error) {
	tx := database.GetDatabaseConnection().Begin()

	DriverTransactionHistory := models.DriverTransactionHistory{
		Remark:        "rejected",
		Status:        "rejected",
		DriverID:      p.DriverID,
		DriverTopupID: p.DriverTopupID,
	}

	if DriverTransactionHistoryErr := tx.Model(&models.DriverTransactionHistory{}).Where("driver_topup_id", p.DriverTopupID).Updates(&DriverTransactionHistory); DriverTransactionHistoryErr.Error != nil {
		tx.Rollback()
		return errors.New("failed to approve driver topup")
	}

	tx.Commit()
	return err
}

func (module *module) RetrieveDriverTopupByUserID(userId uuid.UUID) (m []master.CarManagement, err error) {
	if m, err = module.db.carManagementModel.GetAllCarManagementByDriverID(userId); err != nil {
		return nil, errors.New("failed to transaction history")
	}
	return m, nil
}
