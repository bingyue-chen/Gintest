package integrations

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	jsoniter "github.com/json-iterator/go"

	"Gintest/factories"
	"Gintest/mocks/repositories"
	"Gintest/src/controllers"
	"Gintest/src/entities"
	"Gintest/src/handlers"
	"Gintest/src/middlewares"
)

func TestRetriveProfileSuccess(t *testing.T) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	user := factories.GetUser()

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUserRepository := repositories.NewMockUserRepositoryContract(ctrl)

	mockUserRepository.EXPECT().Find(user.ID).Return(&user, nil)

	provider := controllers.GetNewProvider()

	provider.SetInstance("UserRepository", mockUserRepository)

	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(response)

	context.Params = append(context.Params, gin.Param{Key: "userId", Value: "1"})

	middlewares.BindEntities()(context)
	handlers.RetriveProfile(context, provider.GetProfileService())
	middlewares.ErrorsHandler("true")(context)

	assert.Equal(t, 200, response.Code)

	responseData := entities.ResponseData{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, "success", responseData.Status)

	actualUser := responseData.Data["user"].(map[string]interface{})

	assert.Equal(t, user.ID, int64(actualUser["id"].(float64)))
	assert.Equal(t, user.Email, actualUser["email"])
}

func TestUpdateProfileSuccess(t *testing.T) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	user := factories.GetUser()
	updatedUser := factories.GetUser()

	userUpdation := factories.GetUserUpdation()

	updatedUser.FirstName = "fake first name"
	updatedUser.LastName = "fake last name"

	userUpdation.FirstName = updatedUser.FirstName
	userUpdation.LastName = updatedUser.LastName

	updatedData := map[string]interface{}{}
	updatedData["first_name"] = userUpdation.FirstName
	updatedData["last_name"] = userUpdation.LastName

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUserRepository := repositories.NewMockUserRepositoryContract(ctrl)

	mockUserRepository.EXPECT().Find(user.ID).Return(&user, nil)
	mockUserRepository.EXPECT().Update(&user, updatedData).Return(&updatedUser, nil)

	provider := controllers.GetNewProvider()

	provider.SetInstance("UserRepository", mockUserRepository)

	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(response)

	body, _ := json.Marshal(&userUpdation)
	bodyReader := strings.NewReader(string(body))

	httpRequest := httptest.NewRequest("PATCH", "/api/v1/users/1", bodyReader)
	context.Request = httpRequest

	context.Params = append(context.Params, gin.Param{Key: "userId", Value: "1"})

	middlewares.BindEntities()(context)
	handlers.UpdateProfile(context, provider.GetProfileService())
	middlewares.ErrorsHandler("true")(context)

	assert.Equal(t, 200, response.Code)

	responseData := entities.ResponseData{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, "success", responseData.Status)

	actualUser := responseData.Data["user"].(map[string]interface{})

	assert.Equal(t, updatedUser.ID, int64(actualUser["id"].(float64)))
	assert.Equal(t, updatedUser.Email, actualUser["email"])
	assert.Equal(t, updatedUser.FirstName, actualUser["first_name"])
	assert.Equal(t, updatedUser.LastName, actualUser["last_name"])
}
