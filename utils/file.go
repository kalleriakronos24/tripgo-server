package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/khaimal-group/dto"
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
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
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

func SaveFileToDockerVolume(c *gin.Context, folderName string, documentType string, file *multipart.FileHeader) (outputPath string, err error) {

	workdir, err := os.Getwd()

	filePath := fmt.Sprintf("../files-uploaded/%s/%s", folderName, documentType)
	if file != nil {
		err = c.SaveUploadedFile(file, filepath.Join(workdir, filePath, file.Filename))
		if err != nil {
			return "", err
		}
	}

	document := documentType
	if document == "car-management" {
		outputPath := fmt.Sprintf("storage/po-%s.pdf", RandStringBytes())
		return outputPath, nil
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Message: "Failed to save uploaded files into our server.", Error: err})
		return "", err
	}

	return "", err
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
