package model

import "time"

type BaseModel struct {
	Id        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id" json:"-"`
	CreatedAt time.Time `gorm:"Column:createdAt" json:"-"`
	UpdatedAt time.Time `gorm:"Column:updatedAt" json:"-"`
	DeletedAt time.Time `gorm:"Column:deletedAt" json:"-" sql:"index"`
}
