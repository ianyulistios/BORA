package rabbitmq

import (
	"github.com/streadway/amqp"
)

func SingleQueue(ch *Channel, exchangeName, routingKey, queueName string, bindQueueArgs amqp.Table) (amqp.Queue, error) {
	arg := make(amqp.Table)
	var err error
	var q amqp.Queue
	if q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		arg,       // arguments
	); err != nil {
		failOnError(err)
	} else {
		if err = SingleQueueBind(ch, q, routingKey, exchangeName, bindQueueArgs); err != nil {
			Print(err)
		}
	}
	return q, err
}

func SingleQueueBind(ch *Channel, q amqp.Queue, routingKey, exchangeName string, args amqp.Table) error {
	var err error
	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchangeName,
		false,
		args,
	)
	failOnError(err)
	return err
}

func SingleExchange(ch *Channel, exchangeName string, exchangeType string) error {
	var err error
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

//NewQueue :
func NewQueue(ch *Channel, exchangeName, exchangeNameRetry, queueName, queueNameRetry, exchangeType string, interval int) (amqp.Queue, error) {
	var err error
	var queue amqp.Queue
	if queue, err = DeclareQueue(ch, queueName, queueNameRetry, exchangeName, exchangeNameRetry, interval); err != nil {
		Print(err)
	} else {
		if err = BindQueue(ch, queue, queueNameRetry, exchangeName, exchangeNameRetry); err != nil {
			Print(err)
		}
	}
	return queue, err
}

// DeclareQueue :
func DeclareQueue(ch *Channel, queue, queueRetry, exchangeName, exchangeNameRetry string, interval int) (amqp.Queue, error) {
	arg := make(amqp.Table)
	arg["x-dead-letter-exchange"] = exchangeNameRetry
	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		arg,   // arguments
	)
	failOnError(err)
	// queue retry
	argretry := make(amqp.Table)
	argretry["x-dead-letter-exchange"] = exchangeName
	argretry["x-message-ttl"] = interval
	_, err = ch.QueueDeclare(
		queueRetry, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		argretry,   // arguments
	)
	failOnError(err)
	return q, err
}

//BindQueue :
func BindQueue(ch *Channel, q amqp.Queue, queueRetry, exchangeName, exchangeNameRetry string) error {
	var err error
	err = ch.QueueBind(
		q.Name,
		"",
		exchangeName,
		false,
		nil,
	)
	failOnError(err)
	err = ch.QueueBind(
		queueRetry,
		"",
		exchangeNameRetry,
		false,
		nil,
	)
	failOnError(err)
	return err
}

//DeclareExchange :
func DeclareExchange(ch *Channel, exchangeName string, exchangeNameRetry string, exchangeType string) error {
	var err error
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	err = ch.ExchangeDeclare(
		exchangeNameRetry,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

//ReadMessage :
func ReadMessage(ch *Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err)
	return msgs, err
}

func failOnError(err error) {
	if err != nil {
		Print(err)
	}
}
