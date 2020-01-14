package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/spf13/cobra"
)

// nsqCmd represents the nsq command
var nsqCmd = &cobra.Command{
	Use:   "nsq",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nsq called")
		for i := 0; i < 10; i++ {
			sendMessage()
		}
		time.Sleep(time.Second * 1)
		testNSQ()
	},
}

func init() {
	rootCmd.AddCommand(nsqCmd)
}

func sendMessage() {
	url := "localhost:4150"
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	err = producer.Publish("zgo", []byte("hello world"))
	if err != nil {
		panic(err)
	}
	producer.Stop()
}

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func testNSQ() {
	url := "localhost:4150"

	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("zgo", "struggle", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(url)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}
	}()

	waiter.Wait()
}
