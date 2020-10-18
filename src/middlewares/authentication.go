package middlewares

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"Gintest/src/entities"
	"Gintest/src/utilities"
)

var jwtSecret = []byte(utilities.Getenv("JWT_SECRET"))

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorizationHeader := c.GetHeader("Authorization")
		authorization := strings.Split(authorizationHeader, "Bearer ")

		if len(authorization) < 2 {
			unauthorizedError := entities.UnauthorizedError{Message: "Missing authorization."}
			c.Error(&unauthorizedError)
			c.Abort()
			return
		}

		token := authorization[1]

		tokenClaims, err := jwt.ParseWithClaims(token, &entities.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
			return jwtSecret, nil
		})

		if err != nil {
			var message string
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					message = "Token is malformed"
				} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
					message = "Token could not be verified because of signing problems"
				} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
					message = "Signature validation failed"
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					message = "Token is expired"
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					message = "Token is not yet valid before sometime"
				} else {
					message = "Can not handle this token"
				}
			}

			unauthorizedError := entities.UnauthorizedError{Message: message}
			c.Error(&unauthorizedError)
			c.Abort()
			return
		}

		claims, ok := tokenClaims.Claims.(*entities.Claims)

		if !ok || !tokenClaims.Valid {
			unauthorizedError := entities.UnauthorizedError{Message: "Unvalid."}
			c.Error(&unauthorizedError)
			c.Abort()
			return
		}

		id := c.Param("userId")

		if id != claims.UserId {
			unauthorizedError := entities.UnauthorizedError{Message: "Forbidden."}
			c.Error(&unauthorizedError)
			c.Abort()
			return
		}

		c.Next()
	}

}
