package service

import (
	"zgo/model"
	"zgo/repo"
)

type UserService struct {
	UserRepo repo.UserRepo
}

func CreateUserService(p repo.UserRepo) UserService {
	return UserService{UserRepo: p}
}

func (p *UserService) FindAll() []model.User {
	return p.UserRepo.FindAll()
}

func (p *UserService) FindByID(id uint) model.User {
	return p.UserRepo.FindByID(id)
}

func (p *UserService) Save(user model.User) model.User {
	p.UserRepo.Save(user)

	return user
}

func (p *UserService) Delete(user model.User) {
	p.UserRepo.Delete(user)
}
