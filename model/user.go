package model

type UserModel struct {
	BaseModel
	Username string `gorm:"Column:username;not null" json:"username" binding:"required" validate:"min=1,max=32"`
	Password string `gorm:"Column:password;not null" json:"password" binding:"required" validate:"min=32,max=128"`
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

func Update(u *UserModel) (int, error) {
	DB.Self.Save()
}


