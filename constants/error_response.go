package constants

import (
	"fmt"
	"gitlab.com/odma1/odma-be/dto"
)

func GetErrorResponse(kind string, err error, modelName string) dto.Response {
	var response dto.Response

	switch kind {

	// ==== url params or body ====
	case "payload-error":
		response = dto.Response{
			Data:    false,
			Kind:    "payload-error",
			Message: "Unknown Payload. Please try again later",
			Error:   err.Error(),
		}
		return response
	case "param-query-error":
		response = dto.Response{
			Data:    false,
			Kind:    "param-query-error",
			Message: "Unknown Param. Please try again later",
			Error:   err.Error(),
		}
		return response
	case "uuid-error":
		response = dto.Response{
			Data:    false,
			Kind:    "uuid-error",
			Message: "Unknown UUID. Please try again later.",
			Error:   err.Error(),
		}
		return response

		// ==== data checking ====
	case "data-not-found":
		response = dto.Response{
			Data:    false,
			Kind:    "data-not-found",
			Message: "Data is not found",
		}
		return response
	case "data-existing":
		response = dto.Response{
			Data:    false,
			Kind:    "data-existing",
			Message: "Data is already exist",
		}
		return response

		// ==== common crud operations fails ====
	case "insert-failed":
		response = dto.Response{
			Data:    false,
			Message: fmt.Sprintf("Failed inserting new %s", modelName),
			Kind:    "insert-failed",
			Error:   err.Error(),
		}
		return response
	case "retrieve-failed":
		response = dto.Response{
			Data:    false,
			Kind:    "retrieve-failed",
			Message: fmt.Sprintf("Failed when finding %s", modelName),
			Error:   err.Error(),
		}
		return response
	case "update-failed":
		response = dto.Response{
			Data:    false,
			Kind:    "update-failed",
			Message: fmt.Sprintf("Failed when updating %s", modelName),
			Error:   err.Error(),
		}
		return response
	case "delete-failed":
		response = dto.Response{
			Data:    false,
			Kind:    "delete-failed",
			Message: fmt.Sprintf("Failed when deleting %s", modelName),
			Error:   err.Error(),
		}
		return response

		// ==== other ====
	default:
		return response
	}
}
