package model

import (
	"github.com/fibbery/go-web/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `gorm:"Column:username;not null" json:"username" binding:"required" validate:"min=1,max=32"`
	Password string `gorm:"Column:password;not null" json:"password" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(u).Error
}

func Delete(id int) error {
	u := UserModel{
		BaseModel: BaseModel{Id: id},
	}
	return DB.Self.Delete(&u).Error
}

func Update(u *UserModel) error {
	return DB.Self.Save(u).Error
}

func GetUser(name string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Unscoped().Where("username = ? ", name).First(u)
	return u, d.Error
}

func ListUsers(offset, limit int) ([]*UserModel, error) {
	users := make([]*UserModel, 10)
	db := DB.Self.Unscoped().Offset(offset).Limit(limit).Order("id desc").Find(&users)
	return users, db.Error
}

func (u *UserModel) Validate() error {
	return validator.New().Struct(u)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
