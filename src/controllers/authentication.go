package controllers

import (
	"Gintest/src/handlers"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	provider := GetProvider()

	handlers.Login(c, provider.GetAuthenticationService())
}
