package user

import (
	"demo01/pkg/gormx"
	"fmt"
)

var Repo = newRepo()

func newRepo() *repo {
	return &repo{}
}

type repo struct {
}

func (a *repo) Get(id string) (*User, error) {
	user := &User{}
	err := gormx.DB.Find(user, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *repo) GetByUsername(username string) (*User, error) {
	user := &User{}
	err := gormx.DB.Take(user, "username=?", username).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *repo) GetByEmail(email string) *User {
	user := &User{}
	err := gormx.DB.Take(user, "email=? ", email).Error
	if err != nil {
		return nil
	}
	return user
}

func (a *repo) Create(user *User) error {
	fmt.Println("Create")
	return gormx.DB.Create(user).Error
}

func (a *repo) UpdateColumn(id string, key string, value interface{}) error {
	return gormx.DB.Model(&User{}).Where("id=?", id).UpdateColumn(key, value).Error
}
