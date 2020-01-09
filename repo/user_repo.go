package repo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"zgo/model"
)

type UserRepo struct {
	DB *gorm.DB
}

func CreateUserRepo(DB *gorm.DB) UserRepo {
	return UserRepo{DB: DB}
}

func (p *UserRepo) FindAll() []model.User {
	var users []model.User
	p.DB.Find(&users)

	return users
}

func (p *UserRepo) FindByID(id uint) model.User {
	var user model.User
	p.DB.First(&user, id)

	return user
}

func (p *UserRepo) Save(user model.User) model.User {
	p.DB.Save(&user)

	return user
}

func (p *UserRepo) Delete(user model.User) {
	p.DB.Delete(&user)
}
