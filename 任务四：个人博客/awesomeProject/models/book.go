package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:woaiwd@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"

	// 配置 GORM 日志器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// 打开数据库连接，使用自定义日志器
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 使用我们定义的日志器
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	DB = db
	DB.AutoMigrate(User{}, Post{}, Comment{})
	var user, user1 User
	//user.Name = "王五"
	//DB.Create(&user)
	DB.First(&user, 1)
	var post, post1 Post
	post.ArticleContent = "文章：你好世界"
	//先把数据都准备好，再进行关联新增
	stPost := DB.Model(&user).Association("Posts").Append(&post)
	if stPost != nil {
		log.Fatal(stPost)
	}
	var comment Comment
	DB.First(&post1, 5)
	comment.PostID = post1.ID
	comment.Content = "这文章还真是不错！"
	DB.Create(&comment)
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//先查posts这张表，预加载这张表中存在哪些关联关系，并且自动映射到post结构体中的关联关系字段一次性返回结构体post中
	//其实正常流程是这句DB.Preload("Posts").Preload("Posts.Comments").First(&user1, 2)
	//只是存在关联关系可以简写
	DB.Preload("Posts.Comments").First(&user1, 1)
	fmt.Println(fmt.Sprintf("用户%s发布的篇文章", user1.Name))
	for _, post := range user1.Posts {
		fmt.Printf("文章 ID=%d: %s\n", post.ID, post.ArticleContent)
		for _, comment := range post.Comments {
			fmt.Printf("评论 ID=%d: %s\n", comment.ID, comment.Content)
		}
	}

}
