package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/khaimal-group/db"
	"github.com/kalleriakronos24/khaimal-group/models/master"
)

func (module *module) RetrieveCustomerWebStatisticByUserID(userId uuid.UUID) (ctx int64, err error) {
	if ctx, err = module.db.bookingTransfer.GetCountByCustomerID(userId); err != nil {
		return ctx, errors.New("failed to get statistic count")
	}
	return
}

func (module *module) RetrieveCustomerWebStatisticActiveByUserID(userId uuid.UUID) (ctx int64, err error) {
	if ctx, err = module.db.bookingTransfer.GetCountActiveByCustomerID(userId); err != nil {
		return ctx, errors.New("failed to get statistic count")
	}
	return
}

func (module *module) RetrieveAllCustomerWebStatistic(id uuid.UUID) (m []master.Credentials, err error) {
	return
}

func (module *module) RetrieveAllCustomerWebStatisticPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}

func (module *module) RetrieveCustomerWebStatistic(id uuid.UUID) (m master.Credentials, err error) {
	return
}
