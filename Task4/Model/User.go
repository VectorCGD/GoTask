package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Password  string
	Email     string
	PostCount uint
	Posts     []*Post
}

func (u *User) Register() bool {
	var sid uint
	db.Model(u).Select("id").Where("name", u.Name).Find(&sid)
	if sid > 0 {
		return false
	}
	db.Create(u)
	return true
}

func (u *User) Exist() bool {
	db.Model(u).Select("id").Where("name", u.Name).Find(&u.ID)
	return u.ID > 0
}

func (u *User) LoginVerify() {
	db.Model(u).Select("id, password").Where("name", u.Name).Find(u)
}

func (u *User) GetPosts() (posts []PostBaseInfo) {
	db.Model(u).Preload("Posts").Where("name", u.Name).Take(u)
	posts = make([]PostBaseInfo, 0, len(u.Posts))
	for _, v := range u.Posts {
		posts = append(posts, v.CreateBaseInfo())
	}
	return
}
