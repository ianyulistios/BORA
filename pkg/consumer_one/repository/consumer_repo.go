package repository

import (
	"fmt"

	"github.com/BORA/helper"
	"github.com/BORA/pkg/consumer_one/model"
)

type ConsumerRepo interface {
	ConsumerOne(request model.RequestConsumerOne) chan helper.GlobalResponse
}

type consumerRepo struct {
}

func InitConsumerRepo() ConsumerRepo {
	return &consumerRepo{}
}

func (r *consumerRepo) ConsumerOne(request model.RequestConsumerOne) chan helper.GlobalResponse {
	output := make(chan helper.GlobalResponse)
	go func() {
		fmt.Println("Consumer ONE Repository Start")
		fmt.Println(request.Data)
		fmt.Println("Consumer ONE Repository End")
	}()
	return output
}
