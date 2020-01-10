package cmd

import (
	"fmt"
	"zgo/api"
	"zgo/model"
	"zgo/repo"
	"zgo/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", dbURL)
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

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		db := initDB()
		defer db.Close()

		userAPI := InitUserAPI(db)

		r := gin.Default()

		r.GET("/users", userAPI.FindAll)
		r.GET("/users/:id", userAPI.FindByID)
		r.POST("/users", userAPI.Create)
		r.PUT("/users/:id", userAPI.Update)
		r.DELETE("/users/:id", userAPI.Delete)

		err := r.Run(port)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&port, "port", "p", "", "port of the http server.")
	httpCmd.Flags().StringVarP(&dbURL, "database", "d", "", "URL of SQL Server.")
}
