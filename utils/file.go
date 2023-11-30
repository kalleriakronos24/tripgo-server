package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/odma1/odma-be/dto"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func ConvertMultipartFileToBase64(c *gin.Context, file *multipart.FileHeader, dst string) (base64 string, err error) {

	if file.Filename == "" {
		return "", errors.New("parameter file is empty")
	}

	fileContent, _ := file.Open()
	bytes, err := io.ReadAll(fileContent)
	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		log.Println("jpeg")
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		log.Println("png")
		base64Encoding += "data:image/png;base64,"
	}
	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	if dst != "" {
		// Upload the file to specific dst.
		err = c.SaveUploadedFile(file, ".")
		if err != nil {
			return "", err
		}
	}
	return base64Encoding, err
}

func SaveFileToDockerVolume(c *gin.Context, ownerType string, documentType string, file *multipart.FileHeader) error {

	workdir, err := os.Getwd()
	// Upload the file to specific dst.

	filePath := fmt.Sprintf("../files-uploaded/%s/%s", ownerType, documentType)
	err = c.SaveUploadedFile(file, filepath.Join(workdir, filePath, file.Filename))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Message: "Failed to save uploaded files into our server.", Error: err})
		return err
	}

	return err
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
