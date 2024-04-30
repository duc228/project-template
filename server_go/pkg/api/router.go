package router

import (
	c "server/pkg/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/healthcheck", c.HealthCheck)

	api := router.Group("/api")
	{

		// Albums endpoints
		albums := api.Group("/albums")
		{
			albums.GET("", c.GetAlbums)
			albums.GET(":id", c.GetAlbumByID)
			albums.POST("", c.PostAlbums)
		}

		// Todos endpoints
		todos := api.Group("/todos")
		{
			todos.GET("", c.GetTodos)
		}

	}

	return router

}
