package models

import (
	"gorm.io/gorm"
)

// 账户信息表
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	PostCount int `gorm:"default:0"` // 文章数量统计

	//  一对多：一个用户可以有多篇文章
	Posts []Post `gorm:"foreignKey:UserID"`
}

func updatePostCommentStatus(tx *gorm.DB, PostID uint) error {
	var count int64
	//不管是新增还是删除都先查有没有关联的评论数据，统一返回post要的结果
	tx.Model(&Comment{}).Where("post_id =?", PostID).Count(&count)
	state := "有评论"
	if count == 0 {
		state = "无评论"
	}
	return tx.Model(&Post{}).Where("id=?", PostID).Update("comment_status", state).Error
}

// 新增钩子
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	return updatePostCommentStatus(tx, c.PostID)
}

// 删除钩子
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	return updatePostCommentStatus(tx, c.PostID)
}

type Post struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint // 外键：指向 User（实现多对一）
	ArticleContent string
	CommentStatus  string `gorm:"default:'有评论'"`

	User     User      // 多对一：这篇文章属于哪个用户
	Comments []Comment `gorm:"foreignKey:PostID"` //一对多：一篇文章可以有多个评论
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count+1")).Error
}

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	PostID  uint // 外键：指向 Post
	Content string

	Post Post // 多对一：这条评论属于哪篇文章
}
