package repository

import (
	"fmt"

	"github.com/BORA/helper"
	"github.com/BORA/pkg/consumer_two/model"
)

type ConsumerRepo interface {
	ConsumerTwo(request model.RequestConsumerOne) chan helper.GlobalResponse
}

type consumerRepo struct {
}

func InitConsumerRepo() ConsumerRepo {
	return &consumerRepo{}
}

func (r *consumerRepo) ConsumerTwo(request model.RequestConsumerOne) chan helper.GlobalResponse {
	output := make(chan helper.GlobalResponse)
	go func() {
		fmt.Println("Consumer TWO Repository Start")
		fmt.Println(request.Data)
		fmt.Println("Consumer TWO Repository End")
	}()
	return output
}
