package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/odma1/odma-be/dto"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{Data: "pong"})
}

func UploadFileMultiple(c *gin.Context) {
	// Multipart form
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Error: "No Form file input found"})
	}
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)
		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, ".")
		if err != nil {
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func UploadFileSingle(c *gin.Context) {
	// Multipart form
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Error: "No Form file input found"})
	}
	workdir, err := os.Getwd()

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, filepath.Join(workdir, "../files-uploaded", file.Filename))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Error: err})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
