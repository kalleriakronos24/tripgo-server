package v1

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/khaimal-group/constants"
	"github.com/kalleriakronos24/khaimal-group/dto"
	"github.com/kalleriakronos24/khaimal-group/utils"
)

// AuthLogin godoc
// @Summary      Method to create company
// @Description  A POST Request to create a company
// @Tags         Company
// @Accept       json
// @Produce      json
// @Success      200 {object}	dto.Response
// @Failure      400 {object}	dto.Response
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param data formData dto.InsertFormCompany true "insert company"
// @Param companyCertificate formData file true "company certificate upload"
// @Router       /company/create [post]
func POSTCreateCompany(c *gin.Context) {
	var err error

	pValidator := &dto.InsertFormCompany{}
	if err = c.Bind(&pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	if err := utils.ValidateHTTPPayload(pValidator); err != nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("payload-error", err, ""))
		return
	}

	fCompanyCertificate, _ := c.FormFile("companyCertificate")
	if fCompanyCertificate == nil {
		c.JSON(http.StatusBadRequest, constants.GetErrorResponse("logical", errors.New("field company certificate is required"), "cannot submit if front company certificate is empty"))
		return
	}

	p := &dto.InsertCompany{
		Name:               pValidator.Name,
		PhoneNumber:        pValidator.PhoneNumber,
		Email:              pValidator.Email,
		CompanyName:        pValidator.CompanyName,
		CompanyAddress:     pValidator.CompanyAddress,
		CompanyCountry:     pValidator.CompanyCountry,
		CompanyNumber:      pValidator.CompanyNumber,
		CompanyCertificate: fCompanyCertificate,
	}

	if len(pValidator.Drivers) > 0 {

		for i := 0; i < len(pValidator.Drivers); i++ {
			fDriverLicensePhoto, _ := c.FormFile("drivers.licensePhoto")
			pDriver := &dto.DriverSignup{
				Name:         pValidator.Drivers[i].Name,
				Email:        pValidator.Drivers[i].Email,
				Phone:        pValidator.Drivers[i].Phone,
				DriverType:   "internal",
				LicensePhoto: fDriverLicensePhoto,
				Password:     utils.RandStringGenerator(8),
			}
			log.Printf("%v, %v", pDriver, p)
		}
	}

	// pDriver := &dto.DriverSignup{
	// 	DriverType: "internal",
	// 	// Email: ,
	// }

	// if err = services.Handler.InsertCarManagement(c, p); err != nil {
	// 	c.JSON(http.StatusBadRequest, constants.GetErrorResponse("insert-failed", err, "car management"))
	// 	return
	// }

	c.JSON(http.StatusOK, dto.Response{Data: false, Message: "success"})
}
