// controllers/user_controller.go
package controllers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"web3-blog-backend/config"
	"web3-blog-backend/models"
	"web3-blog-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ServerErr})
		return
	}

	// 检查用户名或邮箱是否已存在
	var existingUser models.User
	err := config.DB.Where("user_name = ? OR emall = ?", user.UserName, user.Emall).First(&existingUser).Error

	if err == nil {
		//  找到了用户，说明已存在
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.AlreadyRegister})
		return
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ServerErr})
		return
	}
	// 加密密码
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.EncryptError})
		return
	}
	user.Password = hashedPassword

	// 保存用户
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.InserError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": utils.Success("注册成功"), "user": gin.H{
		"id":       user.ID,
		"username": user.UserName,
		"email":    user.Emall,
	}})
}
func Login(c *gin.Context) {
	// 用于记录原始请求数据
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[Login] 读取请求体失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ServerErr})
		return
	}
	log.Printf("[Login] 收到原始请求数据: %s", string(body))

	// 重新设置 body，以便 ShouldBindJSON 能读取
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// 定义登录输入结构体
	var input struct {
		Username string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[Login] 参数绑定失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 清理输入（去空格）
	input.Username = strings.TrimSpace(input.Username)
	input.Password = strings.TrimSpace(input.Password)

	log.Printf("[Login] 处理后的用户名: %q", input.Username)
	log.Printf("[Login] 处理后的密码: %q (长度: %d)", input.Password, len(input.Password))

	// 查询用户
	var user models.User
	if err := config.DB.Where("user_name = ?", input.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("[Login] 用户不存在: %s", input.Username)
			c.JSON(http.StatusUnauthorized, gin.H{"error": utils.PasswordError})
		} else {
			log.Printf("[Login] 数据库查询错误: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": utils.ServerErr})
		}
		return
	}

	log.Printf("[Login] 从数据库读取的用户 ID: %d, 用户名: %s", user.ID, user.UserName)

	// 打印数据库中的密码哈希（关键调试）
	log.Printf("[Login] 数据库中的 password 哈希: %q (长度: %d)", user.Password, len(user.Password))

	// 清理哈希（防止有空格或不可见字符）
	cleanHash := strings.TrimSpace(user.Password)
	cleanPass := strings.TrimSpace(input.Password)

	log.Printf("[Login] 清理后的哈希: %q (长度: %d)", cleanHash, len(cleanHash))
	log.Printf("[Login] 清理后的密码: %q (长度: %d)", cleanPass, len(cleanPass))

	// 直接使用 bcrypt 比对，不经过 utils.CheckPassword
	compareErr := bcrypt.CompareHashAndPassword([]byte(cleanHash), []byte(cleanPass))
	if compareErr != nil {
		log.Printf("[Login] bcrypt 比对失败！错误: %v", compareErr)
		c.JSON(http.StatusUnauthorized, gin.H{"error": utils.PasswordError})
		return
	}

	log.Printf(" 密码验证成功！用户 %s 登录", user.UserName)

	// 生成 token
	token, err := utils.GenerateToken(&user)
	if err != nil {
		log.Printf("[Login] Token 生成失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.LoginErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
