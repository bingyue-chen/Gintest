package middlewares

import (
	"Gintest/src/entities"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorsHandler(debug string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		httpStatusCode := http.StatusInternalServerError
		errorBag := entities.ErrorBag{}

		if errors.Is(err.Err, &entities.BindJSONError{}) {
			httpStatusCode = http.StatusUnprocessableEntity
		} else if errors.Is(err.Err, &entities.BadRequestError{}) {
			httpStatusCode = http.StatusBadRequest
		} else if errors.Is(err.Err, &entities.UnauthorizedError{}) {
			httpStatusCode = http.StatusUnauthorized
		} else if errors.Is(err.Err, &entities.ForbiddenError{}) {
			httpStatusCode = http.StatusForbidden
		}

		if httpStatusCode == http.StatusInternalServerError && debug != "true" {
			errorBag.Message = "Oops!! Something went wrong. Please try again later."
		} else {
			errorBag.Message = err.Error()
		}

		errorBag.Message = err.Error()

		responseBag := entities.ResponseBag{}.NewError(errorBag)

		c.JSON(httpStatusCode, responseBag)
		return
	}

}
