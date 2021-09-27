package usecase

import (
	"github.com/BORA/helper"
	"github.com/BORA/pkg/consumer_one/model"
	"github.com/BORA/pkg/consumer_one/repository"
)

type ConsumerUseCase interface {
	ConsumerOne(request model.RequestConsumerOne) chan helper.GlobalResponse
}

type consumerUseCase struct {
	ConsumerRepo repository.ConsumerRepo
}

func InitConsumerUseCase(repo repository.ConsumerRepo) ConsumerUseCase {
	return &consumerUseCase{
		ConsumerRepo: repo,
	}
}

func (u *consumerUseCase) ConsumerOne(request model.RequestConsumerOne) chan helper.GlobalResponse {
	output := make(chan helper.GlobalResponse)
	go func() {
		result := <-u.ConsumerRepo.ConsumerOne(request)
		output <- result
	}()
	return output
}
