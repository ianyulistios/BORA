package rabbitmq

import (
	"log"

	"github.com/BORA/helper/rabbitmq"
)

// NewConnection :
func NewConnection(connStr string) (*rabbitmq.Connection, error) {
	conn, err := rabbitmq.Dial(connStr)
	if err != nil {
		log.Printf("%s: %s", err.Error(), err)
	}
	return conn, err
}
