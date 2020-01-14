package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// websocketCmd represents the websocket command
var websocketCmd = &cobra.Command{
	Use:   "websocket",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("websocket called")
	},
}

func init() {
	rootCmd.AddCommand(websocketCmd)
}
