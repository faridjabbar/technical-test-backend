package exception

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"technical-test-backend/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DatabaseError struct {
	Contains string
	Message  string
}

func ErrorDatabaseMessage(err error, databaseErrors []DatabaseError) string {
	for _, databaseError := range databaseErrors {
		if strings.Contains(err.Error(), databaseError.Contains) {
			return databaseError.Message
		}
	}

	return "database error"
}

func ErrorHandler(c *gin.Context, err interface{}, databaseErrors []DatabaseError) (fatal bool) {
	if validationError(c, err) {
		return false
	}

	if sendToResponseError(c, err) {
		return false
	}

	if permissionDeniedError(c, err) {
		return false
	}

	if recordNotFoundError(c, err) {
		return false
	}

	if unauthorizedError(c, err) {
		return false
	}

	if databaseError(c, err, databaseErrors) {
		return true
	}

	internalServerError(c, err)
	return true
}

func validationError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: "Bad Request",
			Data:    exception.Error(),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}

func recordNotFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && exception.Error() == "record not found" {
		webResponse := web.WebResponse{
			Success: true,
			Message: "Record not found",
		}

		c.JSON(http.StatusOK, webResponse)
		return true
	}
	return false
}

func sendToResponseError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(*ErrorSendToResponse)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func unauthorizedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && (errors.Is(exception, ErrUnauthorized) || errors.Is(exception, ErrRefreshTokenExpired)) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusUnauthorized, webResponse)
		return true
	}
	return false
}

func permissionDeniedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && errors.Is(exception, ErrPermissionDenied) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusForbidden, webResponse)
		return true
	}
	return false
}

func databaseError(c *gin.Context, err interface{}, databaseErrors []DatabaseError) bool {
	exception, ok := err.(error)
	if ok {
		if strings.Contains(exception.Error(), "Duplicate entry") || strings.Contains(exception.Error(), "already exists") ||
			strings.Contains(exception.Error(), "Cannot add or update a child row") ||
			strings.Contains(exception.Error(), "Cannot delete or update a parent row") {
			webResponse := web.WebResponse{
				Success: false,
				Message: ErrorDatabaseMessage(exception, databaseErrors),
			}

			c.JSON(http.StatusBadRequest, webResponse)
			return true
		}
	}
	return false
}

func internalServerError(c *gin.Context, err interface{}) {
	webResponse := web.WebResponse{
		Success: false,
		Message: "Internal Server Error",
	}
	fmt.Println(err)
	c.JSON(http.StatusInternalServerError, webResponse)
}
