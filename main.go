package main

import (
	"go-crud/controllers"
	"go-crud/initializers"
	"go-crud/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	migrate.MigrateDatabases()
}

func main() {
	
	var endp ="/posts/:id"
	r := gin.Default()
	r.POST("/posts", controllers.CreatePosts)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET(endp, controllers.GetSinglePost)
	r.PUT(endp, controllers.UpdatePost)
	r.DELETE(endp, controllers.DeletePost)
	
	r.Run() 
}