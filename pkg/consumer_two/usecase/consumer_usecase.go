package usecase

import (
	"github.com/BORA/helper"
	"github.com/BORA/pkg/consumer_two/model"
	"github.com/BORA/pkg/consumer_two/repository"
)

type ConsumerUseCase interface {
	ConsumerTwo(request model.RequestConsumerOne) chan helper.GlobalResponse
}

type consumerUseCase struct {
	ConsumerRepo repository.ConsumerRepo
}

func InitConsumerUseCase(repo repository.ConsumerRepo) ConsumerUseCase {
	return &consumerUseCase{
		ConsumerRepo: repo,
	}
}

func (u *consumerUseCase) ConsumerTwo(request model.RequestConsumerOne) chan helper.GlobalResponse {
	output := make(chan helper.GlobalResponse)
	go func() {
		result := <-u.ConsumerRepo.ConsumerTwo(request)
		output <- result
	}()
	return output
}
