package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"zgo/api"
	"zgo/model"
	"zgo/repo"
	"zgo/service"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	return db
}

func InitUserAPI(db *gorm.DB) api.UserAPI {
	userRepo := repo.CreateUserRepo(db)
	userService := service.CreateUserService(userRepo)
	userAPI := api.CreateUserAPI(userService)
	return userAPI
}

func main() {
	db := initDB()
	defer db.Close()

	userAPI := InitUserAPI(db)

	r := gin.Default()

	r.GET("/users", userAPI.FindAll)
	r.GET("/users/:id", userAPI.FindByID)
	r.POST("/users", userAPI.Create)
	r.PUT("/users/:id", userAPI.Update)
	r.DELETE("/users/:id", userAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
