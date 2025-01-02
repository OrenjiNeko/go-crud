package main

import (
	"go-crud/controllers"
	"go-crud/initializers"
	"go-crud/middleware"
	"go-crud/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	migrate.MigrateDatabases()
}

func main() {
	
	r := gin.Default()
		
	//user
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/getuser",middleware.RequireAuth ,controllers.GetUser)

	//post
	var endp ="/posts/:id"
	r.POST("/posts", controllers.CreatePosts)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET(endp, controllers.GetSinglePost)
	r.PUT(endp, controllers.UpdatePost)
	r.DELETE(endp, controllers.DeletePost)

	r.Run() 
}