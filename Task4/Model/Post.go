package model

import (
	"gorm.io/gorm"
)

type PostBaseInfo struct {
	Id    int
	Title string
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

func (p *Post) Publish() error {
	res := db.Create(p)
	return res.Error
}

func (p *Post) CreateBaseInfo() (baseInfo PostBaseInfo) {
	baseInfo.Id = int(p.ID)
	baseInfo.Title = p.Title
	return
}

func (p *Post) LoadData() {
	db.Model(p).Select("*").Where("id", p.ID).Find(p)
}

func (p *Post) LoadWithCommnets(author *string) (bool, error) {
	res := db.Preload("Comments").Where("id", p.ID).Take(p)
	db.Model(&User{}).Select("name").Where("id", p.UserID).Find(author)
	return res.RowsAffected > 0, res.Error
}
func (p *Post) UpdateContent() bool {
	var uid uint
	db.Model(p).Select("user_id").Where("id", p.ID).Find(&uid)
	if uid != p.UserID {
		return false
	}
	db.Model(p).Update("content", p.Content)
	return true
}
func (p *Post) Delete() bool {
	var uid uint
	db.Model(p).Select("user_id").Where("id", p.ID).Find(&uid)
	if uid != p.UserID {
		return false
	}
	db.Delete(p)
	return true
}
