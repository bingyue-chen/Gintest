package routes

import (
	"github.com/gin-gonic/gin"

	"Gintest/src/controllers"
	"Gintest/src/middlewares"
)

type ApiRouter struct {
	Router *gin.Engine
}

func (router *ApiRouter) Setup() {

	api := router.Router.Group("/api")
	{

		v1 := api.Group("/v1")
		{
			user := v1.Group("/users")
			{
				user.POST("/login", controllers.Login)
				user.POST("/signup", controllers.Register)

				user.Use(middlewares.Authenticate()).Use(middlewares.BindEntities())
				{
					user.GET("/:userId", controllers.RetriveProfile)
					user.PATCH("/:userId", controllers.UpdateProfile)
				}
			}
		}
	}
}
