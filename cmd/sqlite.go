package cmd

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var orm *xorm.Engine

type User struct {
	Id   int64
	Name string
}

// sqliteCmd represents the sqlite command
var sqliteCmd = &cobra.Command{
	Use:   "sqlite",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sqlite called")
		var err error
		orm, err = xorm.NewEngine("sqlite3", "./test.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer orm.Close()
		orm.ShowSQL(true)
		err = orm.CreateTables(&User{})
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	cliCmd.AddCommand(sqliteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sqliteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sqliteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
