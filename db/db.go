package database

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
	"log"
	"math"
	"strconv"
	"time"

	"gitlab.com/odma1/odma-be/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	SortDesc   string      `json:"sortDesc,omitempty;query:sortDesc"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"rows"`
}

func GetDatabaseConnection() *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(config.AppConfig.DBUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})

	if con, _ := db.DB(); err != nil {
		log.Println("[INIT] failed connecting to PostgresSQL")
		return nil
	} else {
		con.SetConnMaxLifetime(15 * time.Second)
		con.SetMaxOpenConns(90)
		con.SetMaxIdleConns(20)
	}

	if err != nil {
		log.Println("[INIT] failed connecting to PostgresSQL")
		return nil
	}
	return db
}

func DropUnusedColumns(dst interface{}) {

	db := GetDatabaseConnection()
	stmt := db.Statement
	err := stmt.Parse(dst)

	if err != nil {
		log.Println("[INIT] failed parse column on the database ", err.Error())
		return
	}
	fields := stmt.Schema.Fields
	columns, _ := db.Debug().Migrator().ColumnTypes(dst)

	for i := range columns {
		found := false
		for j := range fields {
			if columns[i].Name() == fields[j].DBName {
				found = true
				break
			}
		}
		if !found {
			err := db.Migrator().DropColumn(dst, columns[i].Name())
			if err != nil {
				log.Println("[INIT] failed drop column on the database ", err.Error())
				return
			}
		}
	}
	con, _ := db.DB()
	_ = con.Close()
}

func Paginator(c *gin.Context, value interface{}, relations []string, pagination *Pagination) func(db *gorm.DB) *gorm.DB {

	db := GetDatabaseConnection()
	qPage := c.Query("page")
	qLimit := c.Query("limit")

	page, _ := strconv.Atoi(qPage)
	limit, _ := strconv.Atoi(qLimit)

	if page <= 0 {
		page = 1
	}

	sort := c.Query("sort")

	if sort == "" {
		sort = "created_at"
	}

	var sortWithDirection string
	if sort != "" {
		direction := c.Query("sortDesc")

		if direction == "" {
			direction = "true"
		}

		if direction != "" {
			if direction == "true" {
				pagination.SortDesc = "true"
				sortWithDirection = sort + " desc"
			} else if direction == "false" {
				pagination.SortDesc = "false"
				sortWithDirection = sort + " asc"
			}
		}
	}

	switch {
	case limit <= 0:
		limit = 20
	}

	var totalRows int64
	*db = *db.Model(&value).Count(&totalRows)
	pagination.TotalRows = totalRows
	pagination.Limit = limit

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	pagination.Sort = sortWithDirection
	offset := (page - 1) * limit

	pagination.Page = page

	return func(db *gorm.DB) *gorm.DB {
		if len(relations) > 0 {
			for idx := range relations {
				*db = *db.Offset(offset).Limit(limit).Order(sortWithDirection).Preload(relations[idx])
			}
		}
		*db = *db.Offset(offset).Limit(limit).Order(sortWithDirection)
		con, _ := db.DB()
		_ = con.Close()
		return db
	}
}

func GetLastDocumentNumber(tx *gorm.DB, model interface{}) (m *interface{}, err error) {

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	startingOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	endingOfMonth := startingOfMonth.AddDate(0, 1, -1)

	tx = tx.Model(&model).Where("created_at >= ? AND created_at <= ?", startingOfMonth, endingOfMonth).Last(&model)

	if tx != nil {
		return &model, errors.New("failed to get last document number")
	}

	return &model, nil

}

func GetLastDocumentInvoiceNumber(documentType string, tx *gorm.DB, model interface{}) (m *interface{}, err error) {

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	startingOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	endingOfMonth := startingOfMonth.AddDate(0, 1, -1)

	tx = tx.Model(&model).Where("type = ? AND created_at >= ? AND created_at <= ?", documentType, startingOfMonth, endingOfMonth).Last(&model)

	if tx != nil {
		return &model, errors.New("failed to get last document number")
	}

	return &model, nil

}
