package model

import "time"

type BaseModel struct {
	Id       int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id" json:"-"`
	CreateAt time.Time `gorm:"Column:createAt" json:"-"`
	UpdateAt time.Time `gorm:"Column:updateAt" json:"-"`
	DeleteAt time.Time `gorm:"Column:deleteAt" json:"-"`
}
