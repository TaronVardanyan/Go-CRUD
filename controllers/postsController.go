package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/TaronVardanyan/Go-CRUD/models"
	"github.com/TaronVardanyan/Go-CRUD/initializers"
)

func PostsCreate(c *gin.Context) {
	//Get Request body
	var body struct {
		Body string
		Title string
	}
	
	c.Bind(&body)
	
	//Create a new Post
	post := models.Post{ Title: body.Title, Body: body.Body }
	result := initializers.DB.Create(&post)
	
	//Return a new created post
	if result.Error != nil {
		c.Status(400)
		return
	}
	
	c.JSON(200, gin.H{"post": post})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	
	//Respond with them
	c.JSON(200, gin.H{"posts": posts})
}

func PostsShow(c *gin.Context) {
	//Get Id of post from URL
	id := c.Param("id")
	
	//Get Single Post by primary key
	var post models.Post
	initializers.DB.First(&post, id)
	
	//Respond with them
	c.JSON(200, gin.H{"post": post})
}

func PostsUpdate(c *gin.Context) {
	//Get Post bu url Id
	id := c.Param("id")
	
	//Get Data of Req body
	var body struct {
		Title string
		Body string
	}
	
	c.Bind(&body)
	
	//Find a post from DB
	var post models.Post
	initializers.DB.First(&post, id)
	
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	
	//Respond with them
	c.JSON(200, gin.H{"post": post})
}

func PostsDelete(c *gin.Context) {
	//Get post id from Url
	id := c.Param("id")
	
	//Delete Post
	initializers.DB.Delete(&models.Post{}, id)
	
	//Respond with them
	c.Status(200)
}