package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"Gintest/src/controllers"
	"Gintest/src/entities"
)

type Binder func(*gin.Context, int64)

func BindEntities() gin.HandlerFunc {
	return func(c *gin.Context) {

		binding_map := map[string]Binder{
			"userId": UserBinder,
		}

		for name, binder := range binding_map {

			if id_string := c.Param(name); id_string != "" {

				id, err := strconv.ParseInt(id_string, 10, 64)

				if err != nil {
					c.Error(err)
					c.Abort()
					return
				}

				binder(c, id)
			}
		}

		c.Next()
	}
}

var UserBinder Binder = func(c *gin.Context, id int64) {

	user, err := controllers.GetProvider().GetUserRepository().Find(id)

	if err != nil {
		err := entities.BadRequestError{err.Error()}
		c.Error(&err)
		c.Abort()
		return
	}

	c.Set("UserId", id)
	c.Set("User", user)
}
