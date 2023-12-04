package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
	"log"
	"math"
	"strconv"

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
	if err != nil {
		log.Println("[INIT] failed connecting to PostgresSQL")
		return nil
	}
	return db
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
		return db
	}
}
