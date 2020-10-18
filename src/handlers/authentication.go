package handlers

import (
	"Gintest/src/entities"
	"Gintest/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary User Login
// @Description Validate email and password, then generate api key.
// @Accept  json
// @Produce  json
// @Param email body string true "User Email"
// @Param password body string true "User Password"
// @Success 200 {object} entities.ResponseBag{status=string,data=object{user=entities.User,token=string,token_type=string}}
// @Failure 400 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 401 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 422 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Failure 500 {object} entities.ResponseBag{status=string,error=entities.ErrorBag}
// @Router /users/login [post]
func Login(c *gin.Context, authenticationService services.AuthenticationServiceContract) {

	var credential entities.UserCredential

	if err := c.ShouldBindJSON(&credential); err != nil {
		bindJsonErr := entities.BindJSONError{}
		bindJsonErr.SetErrors(err)
		c.Error(&bindJsonErr)
		return
	}

	user, ok := authenticationService.Authenticate(&credential)

	if !ok {
		err := entities.BadRequestError{"Wrong credential."}
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
