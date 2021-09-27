package consumer

import (
	"fmt"

	"github.com/BORA/constants"
	"github.com/BORA/helper/rabbitmq"
	"github.com/BORA/pkg/consumer_two/handler"
	"github.com/BORA/pkg/consumer_two/repository"
	"github.com/BORA/pkg/consumer_two/usecase"
)

func ConsumerTwoProcessHandler(rabbitmqConn *rabbitmq.Connection) {

	subCh, err := rabbitmqConn.Channel()
	if err != nil {
		fmt.Printf("Error when initializing new channel (%v)\n", err.Error())
	}

	if err = rabbitmq.SingleExchange(subCh, constants.ExchangeName, constants.ExchangeType); err != nil {
		fmt.Printf("Warning : (MerchantStatusHandler) : Failed to declare exchange (%v)\n", err.Error())
	}

	repo := repository.InitConsumerRepo()
	useCase := usecase.InitConsumerUseCase(repo)
	handler := handler.InitConsumerHandler(subCh, useCase)
	handler.ProcessMessage()
}
