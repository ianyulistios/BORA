package rabbitmq

import "log"

// Debug debug log flag
var Debug bool

func Print(args ...interface{}) {
	if Debug {
		log.Print(args...)
	}
}

func Printf(format string, args ...interface{}) {
	if Debug {
		log.Printf(format, args...)
	}
}
