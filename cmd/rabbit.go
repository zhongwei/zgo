package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// rabbitCmd represents the rabbit command
var rabbitCmd = &cobra.Command{
	Use:   "rabbit",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connecting to RabbitMQ ...")
		conn, _ := amqp.Dial(amqpURL)
		defer conn.Close()

		ch, _ := conn.Channel()
		defer ch.Close()

		q, _ := ch.QueueDeclare(
			"DemoQueue", //name
			true,        //durable
			false,       //delete when unused
			false,       //exclusive
			false,       //no-wait
			nil,         //arguements
		)

		body := "Hello world!"

		//Publish to the queue
		ch.Publish(
			"",     //exchange
			q.Name, //routing key
			false,  //mandatory
			false,  //immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		msgs, _ := ch.Consume(
			"DemoQueue", //queue
			"",          //consumer
			true,        //auto-ack
			false,       //exclusive
			false,       //no-local
			false,       //no-wait
			nil,         //args
		)

		msgCount := 0
		go func() {
			for d := range msgs {
				msgCount++
				fmt.Printf("\nMessage Count: %d, Message Body: %s\n", msgCount, d.Body)
			}
		}()

		select {
		case <-time.After(time.Second * 2):
			fmt.Printf("Total Messages Fetched: %d\n", msgCount)
			fmt.Println("No more messages in queue. Timing out...")
		}
	},
}

func init() {
	rootCmd.AddCommand(rabbitCmd)
	rabbitCmd.Flags().StringVarP(&amqpURL, "amqp", "q", "", "URL of RabbitMQ.")
}
