package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"zgo/api"
	"zgo/repo"
	"zgo/service"
)

func initUserAPI(db *gorm.DB) api.UserAPI {
	wire.Build(repo.CreateUserRepo, service.CreateUserService, api.CreateUserAPI)

	return api.UserAPI{}
}
