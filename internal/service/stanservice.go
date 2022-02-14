package service

import (
	"github.com/nats-io/stan.go"
	"log"
)

func getStan() stan.Conn {

	natsConnection, err := stan.Connect("test-cluster", "me")
	if err != nil {
		log.Fatal(err)
	}

	return natsConnection
}
