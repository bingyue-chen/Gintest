package handlers

import (
	"Gintest/src/entities"
	"Gintest/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Retrie Profile godoc
// @Summary Retrie User's Profile
// @Description Retrie User's Profile
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Param id path int true "User ID"
// @Success 200 {object} entities.ResponseBag{status=string,data=object{user=entities.User}}
// @Failure 400 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 401 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 422 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 500 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Router /users/{id} [get]
func RetriveProfile(c *gin.Context, profileService services.ProfileServiceContract) {

	var user *entities.User

	if val, ok := c.Get("User"); ok && val != nil {
		user, _ = val.(*entities.User)
	}

	responseBag := entities.ResponseBag{}.New(gin.H{"user": user})

	c.JSON(http.StatusOK, responseBag)
}

// Update Profile godoc
// @Summary Update User's Profile
// @Description Update User's Profile
// @Accept  json
// @Produce  json
// @Security JWTAuth
// @Param id path int true "User ID"
// @Param first_name body string false "User First Name"
// @Param last_name body string false "User Last Name"
// @Param password body string false "User Password"
// @Success 200 {object} entities.ResponseBag{status=string,data=object{user=entities.User}}
// @Failure 400 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 401 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 422 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 500 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Router /user/{id} [patch]
func UpdateProfile(c *gin.Context, profileService services.ProfileServiceContract) {

	var user *entities.User

	if val, ok := c.Get("User"); ok && val != nil {
		user, _ = val.(*entities.User)
	}

	var updation entities.UserUpdation

	if err := c.ShouldBindJSON(&updation); err != nil {
		bindJsonErr := entities.BindJSONError{}
		bindJsonErr.SetErrors(err)
		c.Error(&bindJsonErr)
		return
	}

	user, err := profileService.Update(user, &updation)

	if err != nil {
		err := entities.BadRequestError{err.Error()}
		c.Error(&err)
		return
	}

	responseBag := entities.ResponseBag{}.New(gin.H{"user": user})

	c.JSON(http.StatusOK, responseBag)
}
