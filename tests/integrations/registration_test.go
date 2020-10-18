package integrations

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	jsoniter "github.com/json-iterator/go"

	"Gintest/factories"
	"Gintest/mocks/repositories"
	"Gintest/src/controllers"
	"Gintest/src/entities"
	"Gintest/src/handlers"
	"Gintest/src/middlewares"
)

func TestRegisterSuccess(t *testing.T) {

	user := factories.GetUser()
	userRegisteration := factories.GetUserRegisteration()

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUserRepository := repositories.NewMockUserRepositoryContract(ctrl)

	err := gorm.ErrRecordNotFound

	mockUserRepository.EXPECT().FindByEmail(userRegisteration.Email).Return(nil, err)

	mockUserRepository.EXPECT().Create(gomock.Any()).Return(&user, nil)

	provider := controllers.GetNewProvider()

	provider.SetInstance("UserRepository", mockUserRepository)

	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(response)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	body, _ := json.Marshal(&userRegisteration)

	bodyReader := strings.NewReader(string(body))

	httpRequest := httptest.NewRequest("POST", "/api/v1/register", bodyReader)
	context.Request = httpRequest

	handlers.Register(context, provider.GetRegistrationService(), provider.GetAuthenticationService())
	middlewares.ErrorsHandler("true")(context)

	assert.Equal(t, 200, response.Code)

	responseData := entities.ResponseData{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, "success", responseData.Status)

	actualUser := responseData.Data["user"].(map[string]interface{})

	assert.Equal(t, user.ID, int64(actualUser["id"].(float64)))
	assert.Equal(t, user.Email, actualUser["email"])
}
