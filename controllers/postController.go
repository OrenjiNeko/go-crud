package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {
	//get data
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)
	
	//create post
	post := models.Post{Title: body.Title,Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error !=nil{
		c.Status(400)
		return
	}

	//return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context){
	//get data
	var posts []models.Post
	initializers.DB.Find(&posts)

	//return data
	c.JSON(200,gin.H{
		"post":posts,
	})
}

func GetSinglePost(c *gin.Context){
	//get id
	id := c.Param("id")

	//get data
	var post models.Post
	initializers.DB.First(&post,id)

	//return data
	c.JSON(200,gin.H{
		"post":post,
	})
}

func UpdatePost(c *gin.Context){
	// get id from url
	id := c.Param("id")

	//get data 
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	//fint the post to update
	var post models.Post
	initializers.DB.First(&post,id)

	//update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:body.Body,
	})

	//return data
	c.JSON(200,gin.H{
		"post":post,
	})
}

func DeletePost(c *gin.Context){
	//get id from url
	id := c.Param("id")

	//delete the post
	initializers.DB.Delete(&models.Post{},id)

	//respond
	c.JSON(200,gin.H{
		"message":"berhasil delete post",
	})
}
