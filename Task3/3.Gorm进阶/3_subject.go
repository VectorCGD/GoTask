package main

import (
	"gorm.io/gorm"
)

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var postCount uint
	tx.Model(&User{}).Select("post_count").Where("id", p.UserID).First(&postCount)
	tx.Model(&User{}).Where("id", p.UserID).Update("post_count", postCount+1)
	return
}

func PostCreateHook(tx *gorm.DB) {
	tx.Create(&Post{Title: "新文章", Content: "文章内容", CmtState: "无评论", CmtCount: 0, UserID: 2})
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var cmtCount uint
	tx.Model(&Post{}).Select("cmt_count").Where("id", c.PostID).First(&cmtCount)
	cmtCount--
	if cmtCount < 1 {
		tx.Model(&Post{}).Where("id", c.PostID).Updates(map[string]interface{}{"cmt_count": cmtCount, "cmt_state": "无评论"})
	} else {
		tx.Model(&Post{}).Where("id", c.PostID).Update("cmt_count", cmtCount)
	}
	return
}

func CommentDeleteHook(tx *gorm.DB) {
	var cmt Comment
	tx.First(&cmt, 4)
	tx.Delete(&cmt)
}
