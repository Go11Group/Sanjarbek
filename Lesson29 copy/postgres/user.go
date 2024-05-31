package postgres

import (
	"module/model"
	"gorm.io/gorm"
)

type UserRepo struct{
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB)*UserRepo{
	return &UserRepo{Db: db}
}


func(U *UserRepo) CreateTable(user model.User)error{
	err := U.Db.AutoMigrate(&user)
	if err != nil{
		return err
	}
	return nil
}

func(U *UserRepo) Create(user model.User){
	U.Db.Create(&user)
}

func(U *UserRepo) Read()([]model.User){
	var users = []model.User{}
	U.Db.Find(&users)
	return users
}

func(U *UserRepo) Update(user model.User, id int){
	U.Db.Model(&user).Updates(model.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password, Field: user.Field, IsEmployee: user.IsEmployee})
}

func(U *UserRepo) Delete(user model.User){
	U.Db.Delete(&user, 1)
}