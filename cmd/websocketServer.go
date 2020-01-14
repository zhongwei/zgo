package cmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/spf13/cobra"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

// websocketServerCmd represents the websocketServer command
var websocketServerCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Websocket Server called")
		flag.Parse()
		log.SetFlags(0)
		http.HandleFunc("/echo", echo)
		log.Fatal(http.ListenAndServe(*addr, nil))
	},
}

func init() {
	websocketCmd.AddCommand(websocketServerCmd)
}

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
