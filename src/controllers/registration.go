package controllers

import (
	"Gintest/src/handlers"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	provider := GetProvider()

	handlers.Register(c, provider.GetRegistrationService(), provider.GetAuthenticationService())
}
