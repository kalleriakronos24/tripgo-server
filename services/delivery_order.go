package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "gitlab.com/odma1/odma-be/db"
	"gitlab.com/odma1/odma-be/dto"
	"gitlab.com/odma1/odma-be/models"
	"gitlab.com/odma1/odma-be/utils"
	"strconv"
	"strings"
	"time"
)

type CheckExistingDeliveryOrderStruct struct {
	*models.DeliveryOrder
}

func (module *module) RetrieveDeliveryOrder(id uuid.UUID) (m models.DeliveryOrder, err error) {
	if m, err = module.db.deliveryOrderModel.GetOneDeliveryOrderByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllDeliveryOrder(id uuid.UUID) (m []models.DeliveryOrder, err error) {
	if m, err = module.db.deliveryOrderModel.GetAllDeliveryOrder(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllDeliveryOrderPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.deliveryOrderModel.GetAllDeliveryOrderPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertDeliveryOrder(c *gin.Context, p *dto.InsertDeliveryOrder) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	DeliveryOrderModel := models.DeliveryOrder{}

	_, deliveryLastDataErr := database.GetLastDocumentNumber(tx, &DeliveryOrderModel)

	var formattedNumber string
	var finalNumber string
	if deliveryLastDataErr != nil {
		now := time.Now()
		year := now.Year()
		month := now.Month()
		monthInRoman := utils.IntegerToRoman(int(month))
		formattedNumber = fmt.Sprintf("%s/%s/%s/%d", "0001", "DO", monthInRoman, year)
		finalNumber = "0001"
	} else {
		parsedNumber, _ := strconv.Atoi(DeliveryOrderModel.Sequence)
		strLastNumber := strconv.Itoa(parsedNumber)
		getFrontDigitNumber := strings.Replace(DeliveryOrderModel.Sequence, strLastNumber, "", -1)
		convertActiveNumberToStr, _ := strconv.Atoi(strLastNumber)
		addActiveNumberByOne := convertActiveNumberToStr + 1
		convertAddedActiveNumberToStr := strconv.Itoa(addActiveNumberByOne)
		poNumber := getFrontDigitNumber + convertAddedActiveNumberToStr

		now := time.Now()
		year := now.Year()
		month := now.Month()
		monthInRoman := utils.IntegerToRoman(int(month))
		formattedNumber = fmt.Sprintf("%s/%s/%s/%d", poNumber, "DO", monthInRoman, year)
		finalNumber = poNumber
	}

	DeliveryOrder := models.DeliveryOrder{
		Number:                 formattedNumber,
		ContactPerson:          p.ContactPerson,
		PhoneNumber:            p.PhoneNumber,
		Address:                p.Address,
		Note:                   p.Note,
		Date:                   p.Date,
		Status:                 p.Status,
		Sequence:               finalNumber,
		OperatingActivityID:    p.OperatingActivityID,
		DeliveryOrderCreatedBy: p.CreatedBy,
	}

	if DeliveryOrderErr := tx.Create(&DeliveryOrder); err != nil {
		tx.Rollback()
		return DeliveryOrderErr.Error
	}
	tx.Commit()
	return
}

func (module *module) UpdateDeliveryOrder(c *gin.Context, id uuid.UUID, p *dto.UpdateDeliveryOrder) (err error) {

	tx := database.GetDatabaseConnection().Begin()

	DeliveryOrder := models.DeliveryOrder{
		ID: id,
		//Number:                 p.Number,
		ContactPerson:          p.ContactPerson,
		PhoneNumber:            p.PhoneNumber,
		Address:                p.Address,
		Note:                   p.Note,
		Date:                   p.Date,
		Status:                 p.Status,
		OperatingActivityID:    p.OperatingActivityID,
		DeliveryOrderUpdatedBy: p.UpdatedBy,
	}

	if DeliveryOrderErr := tx.Updates(&DeliveryOrder); err != nil {
		tx.Rollback()
		return DeliveryOrderErr.Error
	}

	tx.Commit()
	return
}

func (module *module) CheckExistingDeliveryOrder(id string, param CheckExistingDeliveryOrderStruct) (err error) {

	if param.Number != "" {
		if _, dbErr := module.db.deliveryOrderModel.GetOneDeliveryOrderByNumber(param.Number); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if param.OperatingActivityID.String() != "" {
		if _, dbErr := module.db.deliveryOrderModel.GetOneDeliveryOrderByOperatingActivityID(param.OperatingActivityID); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.deliveryOrderModel.GetOneDeliveryOrderByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
