package controllers

import (
	"Gintest/src/handlers"
	"github.com/gin-gonic/gin"
)

func RetriveProfile(c *gin.Context) {

	provider := GetProvider()

	handlers.RetriveProfile(c, provider.GetProfileService())
}

func UpdateProfile(c *gin.Context) {

	provider := GetProvider()

	handlers.UpdateProfile(c, provider.GetProfileService())
}
