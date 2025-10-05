package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Password  string
	Email     string
	PostCount uint
	Posts     []*Post
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	CmtState string
	CmtCount uint
	UserID   uint
	Comments []*Comment
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
	User    *User
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/youbase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	//2_subject
	GetPostsByUserName(db, "月尽")
	commentsCountMax(db)

	//3subject
	PostCreateHook(db)
	CommentDeleteHook(db)
}
