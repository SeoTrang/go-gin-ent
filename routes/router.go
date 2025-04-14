package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/seotrang/go-ent/controllers"
	"github.com/seotrang/go-ent/ent"
)

func SetupRoutes(r *gin.Engine, client *ent.Client) {
	api := r.Group("/api")
	{
		api.GET("/ping", controllers.Ping)
		api.GET("/users", func(c *gin.Context) {
			controllers.GetUsers(c, client)
		})
		api.GET("/users/:id", func(c *gin.Context) {
			controllers.GetUserByID(c, client)
		})
		api.POST("/users", func(c *gin.Context) {
			controllers.CreateUser(c, client)
		})
		api.PUT("/users/:id", func(c *gin.Context) {
			controllers.UpdateUser(c, client)
		})
		api.DELETE("/users/:id", func(c *gin.Context) {
			controllers.DeleteUser(c, client)
		})
	}
}
