package main

import (
	config "phonebook/config"
	phoneCon "phonebook/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.CheckConfiguration()
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	version := router.Group("/api/v1")
	{
		phones := version.Group("/phones")
		{
			phones.POST("/", phoneCon.PostData)
			phones.GET("/", phoneCon.GetData)
			phones.GET("/:id", phoneCon.GetOne)
			phones.PUT("/:id", phoneCon.PutOne)
			phones.DELETE("/:id", phoneCon.DeleteOne)
		}
	}
	return router
}
