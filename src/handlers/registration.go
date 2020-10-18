package handlers

import (
	"Gintest/src/entities"
	"Gintest/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register godoc
// @Summary User Registration
// @Description Create new user with email, password and name.
// @Accept  json
// @Produce  json
// @Param email body string true "User Email"
// @Param password body string true "User Password"
// @Param name body string true "User Name"
// @Success 200 {object} entities.ResponseBag{status=string,data=object{user=entities.User,token=string,token_type=string}}
// @Failure 400 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 401 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 422 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 500 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Router /users/singup [post]
func Register(c *gin.Context, registrationService services.RegistrationServiceContract, authenticationService services.AuthenticationServiceContract) {

	userRegisteration := entities.UserRegisteration{}

	if err := c.ShouldBindJSON(&userRegisteration); err != nil {
		bindJsonErr := entities.BindJSONError{}
		bindJsonErr.SetErrors(err)
		c.Error(&bindJsonErr)
		return
	}

	user, err := registrationService.Register(&userRegisteration)

	if err != nil {
		err := entities.BadRequestError{err.Error()}
		c.Error(&err)
		return
	}

	token, err := authenticationService.GenerateToken(user)
	if err != nil {
		c.Error(err)
		return
	}

	responseBag := entities.ResponseBag{}.New(gin.H{"user": user, "token": token, "token_type": "Bearer"})

	c.JSON(http.StatusOK, responseBag)

}
