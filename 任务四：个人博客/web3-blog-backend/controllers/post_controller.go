package controllers

import (
	"net/http"
	"web3-blog-backend/config"
	"web3-blog-backend/models"
	"web3-blog-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	//声明文章创建结构体
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	//获取文章信息
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ResponseErro})
		return
	}
	//获取当前登录的用户信息
	user, _ := c.Get("currentUser")
	currentUser := user.(*models.User)
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  currentUser.ID,
	}
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ResponseErro_1})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"post": utils.Success(post)})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	//将user下关联的评论所有信息都查出来放入post
	if err := config.DB.Preload("User").Preload("Comments").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": utils.SelectError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": utils.Success(posts)})
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("ID")
	if err := config.DB.Preload("User").Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": utils.SelectError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": utils.Success(post)})
}
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post *models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": utils.SelectError})
		return
	}
	//获取当前用户
	user, _ := c.Get("currentUser")
	currentUser := user.(models.User)
	if post.ID != currentUser.ID {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": utils.ResponseErro_2})
		return
	}
	var input struct {
		Title   *string `json:"title"`
		Content *string `json:"content"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": utils.ResponseErro_3})
		return
	}
	if input.Title != nil {
		post.Title = *input.Title
	}
	if input.Content != nil {
		post.Content = *input.Content
	}
	if err := config.DB.Save("Post").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": utils.ResponseErro_4})
		return
	}
	c.JSON(200, utils.Success("操作成功"))
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	user, _ := c.Get("currentUser")
	currentUser := user.(models.User)

	if post.UserID != currentUser.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权删除此文章"})
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
