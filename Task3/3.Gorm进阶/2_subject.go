package main

import (
	"fmt"

	"gorm.io/gorm"
)

func GetPostsByUserName(db *gorm.DB, name string) {
	tar := User{}
	db.Preload("Posts").Preload("Posts.Comments").Where("name = ?", name).Take(&tar)
	for i := 0; i < len(tar.Posts); i++ {
		for c := 0; c < len(tar.Posts[i].Comments); c++ {
			db.Preload("User").Take(tar.Posts[i].Comments[c])
		}
	}
	//output
	for _, v := range tar.Posts {
		fmt.Println(*v)
		for _, c := range v.Comments {
			fmt.Println(*c)
		}
	}
}

func commentsCountMax(db *gorm.DB) {

	var allPost []Post
	db.Preload("Comments").Find(&allPost)
	maxPosts := make([]Post, 0, 4)
	for _, v := range allPost {
		if len(maxPosts) < 1 {
			maxPosts = append(maxPosts, v)
			continue
		}
		if len(v.Comments) > len(maxPosts[0].Comments) {
			maxPosts = maxPosts[:0]
			maxPosts = append(maxPosts, v)
		} else if len(v.Comments) == len(maxPosts[0].Comments) {
			maxPosts = append(maxPosts, v)
		}
	}

	//output
	for _, v := range maxPosts {
		fmt.Println(v)
		for _, c := range v.Comments {
			fmt.Println(*c)
		}
	}
}
