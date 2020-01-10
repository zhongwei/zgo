package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
		if err != nil {
			fmt.Println(err)
		}

		err = client.Set("id1234", json, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		val, err := client.Get("id1234").Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}
