package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/elastic/go-elasticsearch/v7"
)

// elasticCmd represents the elastic command
var elasticCmd = &cobra.Command{
	Use:   "elastic",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		es, _ := elasticsearch.NewDefaultClient()
		fmt.Println(es.Info())
	},
}

func init() {
	rootCmd.AddCommand(elasticCmd)
}
