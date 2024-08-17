package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/khaimal-group/constants"
	"github.com/kalleriakronos24/khaimal-group/dto"
)

type AuthResponseData struct {
	message string
}

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{Error: "User not authenticated", Data: false})
	}
}
