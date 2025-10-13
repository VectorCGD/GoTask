package model

import "gorm.io/gorm"

type CmtInfo struct {
	Cor     string
	Content string
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
	User    *User
}

func (c *Comment) CreateInfor() CmtInfo {
	var uname string
	res := db.Model(c.User).Select("name").Where("id", c.UserID).Find(&uname)
	if res.Error != nil || res.RowsAffected == 0 {
		uname = "0"
	}
	return CmtInfo{Cor: uname, Content: c.Content}
}
func (c *Comment) Publish() {
	db.Create(c)
}
