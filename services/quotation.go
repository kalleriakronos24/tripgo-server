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

type CheckExistingQuotationStruct struct {
	*models.Quotation
}

func (module *module) RetrieveAllQuotation(id uuid.UUID) (m []models.Quotation, err error) {
	if m, err = module.db.quotationModel.GetAllQuotation(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveAllQuotationPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	if pagination, err = module.db.quotationModel.GetAllQuotationPaginated(c, id); err != nil {
		return pagination, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) RetrieveQuotation(id uuid.UUID) (m models.Quotation, err error) {
	if m, err = module.db.quotationModel.GetOneQuotationByID(id); err != nil {
		return m, fmt.Errorf(err.Error())
	}
	return
}

func (module *module) InsertQuotation(p *dto.InsertQuotation) (err error) {

	tx := database.GetDatabaseConnection()
	QuotationModel := models.Quotation{}

	_, quotationLastDataErr := database.GetLastDocumentNumber(tx, &QuotationModel)

	if quotationLastDataErr != nil {
		QuotationModel.Sequence = "0001"
	}

	parsedNumber, _ := strconv.Atoi(QuotationModel.Sequence)
	strLastNumber := strconv.Itoa(parsedNumber)
	getFrontDigitNumber := strings.Replace(QuotationModel.Sequence, strLastNumber, "", -1)
	convertActiveNumberToStr, _ := strconv.Atoi(strLastNumber)
	addActiveNumberByOne := convertActiveNumberToStr + 1
	convertAddedActiveNumberToStr := strconv.Itoa(addActiveNumberByOne)
	finalNumber := getFrontDigitNumber + convertAddedActiveNumberToStr

	now := time.Now()
	year := now.Year()
	month := now.Month()
	monthInRoman := utils.IntegerToRoman(int(month))
	formattedNumber := fmt.Sprintf("%s/%s/%s/%d", finalNumber, "SPH", monthInRoman, year)

	if p.OperatingActivityID.String() != "" {
		Quotation := models.Quotation{
			Number:              formattedNumber,
			FrancoArea:          p.FrancoArea,
			PaymentTerm:         p.PaymentTerm,
			SendAfter:           p.SendAfter,
			Date:                p.Date,
			Sequence:            finalNumber,
			OperatingActivityID: p.OperatingActivityID,
			QuotationCreatedBy:  p.CreatedBy,
		}

		if QuotationErr := tx.Create(&Quotation); err != nil {
			tx.Rollback()
			return QuotationErr.Error
		}
	} else {
		Quotation := models.Quotation{
			Number:             formattedNumber,
			FrancoArea:         p.FrancoArea,
			PaymentTerm:        p.PaymentTerm,
			SendAfter:          p.SendAfter,
			Date:               p.Date,
			Sequence:           finalNumber,
			QuotationCreatedBy: p.CreatedBy,
		}

		if QuotationErr := tx.Create(&Quotation); err != nil {
			tx.Rollback()
			return QuotationErr.Error
		}
	}
	tx.Commit()
	return
}

func (module *module) UpdateQuotation(id uuid.UUID, p *dto.UpdateQuotation) (err error) {

	if p.OperatingActivityID.String() != "" {
		if err = module.db.quotationModel.UpdateQuotation(id, models.Quotation{
			FrancoArea:          p.FrancoArea,
			PaymentTerm:         p.PaymentTerm,
			SendAfter:           p.SendAfter,
			Date:                p.Date,
			OperatingActivityID: p.OperatingActivityID,
			QuotationUpdatedBy:  p.UpdatedBy,
		}); err != nil {
			return errors.New(err.Error())
		}
	} else {
		if err = module.db.quotationModel.UpdateQuotation(id, models.Quotation{
			FrancoArea:         p.FrancoArea,
			PaymentTerm:        p.PaymentTerm,
			SendAfter:          p.SendAfter,
			Date:               p.Date,
			QuotationUpdatedBy: p.UpdatedBy,
		}); err != nil {
			return errors.New(err.Error())
		}
	}
	return
}

func (module *module) CheckExistingQuotation(id string, param CheckExistingQuotationStruct) (err error) {

	if param.Number != "" {
		if _, dbErr := module.db.quotationModel.GetOneQuotationByNumber(param.Number); dbErr != nil {
			return errors.New(dbErr.Error())
		}
		return
	}

	if id != "" {
		uid, _ := uuid.Parse(id)
		if _, dbErr := module.db.quotationModel.GetOneQuotationByID(uid); dbErr != nil {
			return errors.New(dbErr.Error())
		}
	}
	return
}
