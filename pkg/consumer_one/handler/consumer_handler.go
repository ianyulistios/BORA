package handler

import (
	"encoding/json"
	"fmt"

	"github.com/BORA/constants"
	"github.com/BORA/helper/rabbitmq"
	"github.com/BORA/pkg/consumer_one/model"
	"github.com/BORA/pkg/consumer_one/usecase"
	"github.com/streadway/amqp"
)

type ConsumerHandler interface {
	ProcessMessage()
}

type consumerHandler struct {
	SubConn         *rabbitmq.Channel
	ConsumerUseCase usecase.ConsumerUseCase
}

func InitConsumerHandler(subConn *rabbitmq.Channel, usecase usecase.ConsumerUseCase) ConsumerHandler {
	return &consumerHandler{
		SubConn:         subConn,
		ConsumerUseCase: usecase,
	}
}

func (h *consumerHandler) ProcessMessage() {
	args := make(amqp.Table)
	var message <-chan amqp.Delivery
	args["format"] = "pdf"
	args["type"] = "report"
	args["x-match"] = "any"
	queue, err := rabbitmq.SingleQueue(h.SubConn, constants.ExchangeName, constants.RoutingKeyOne, constants.QueueNameOne, args)
	if message, err = rabbitmq.ReadMessage(h.SubConn, queue); err != nil {
		fmt.Println(err)
	} else {
		for msg := range message {
			var msgBody model.RequestConsumerOne
			if err := json.Unmarshal(msg.Body, &msgBody); err != nil {
				fmt.Println("Error when getting message body" + err.Error())
			}

			h.ConsumerUseCase.ConsumerOne(msgBody)

			msg.Ack(true)
			msg.Reject(false)
		}
	}
}
