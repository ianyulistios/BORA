package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/BORA/cmd/consumer"
	"github.com/BORA/config/rabbitmq"
	"github.com/BORA/constants"
)

func main() {
	rabbitMQConn, err := rabbitmq.NewConnection(constants.RabbitURL)

	if err != nil {
		fmt.Printf("Error initiating rabbitmq publisher connection (%v)\n", err.Error())
		os.Exit(2)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("starting handler one")
		consumer.ConsumerOneProcessHandler(rabbitMQConn)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("starting handler two")
		consumer.ConsumerTwoProcessHandler(rabbitMQConn)
	}()

	wg.Wait()
	fmt.Printf("APPLICATION EXIT")
}
