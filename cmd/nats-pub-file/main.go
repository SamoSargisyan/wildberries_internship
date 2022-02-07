package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"l0/internal/domain"
	"log"
	"os"
)

const (
	clusterID = "test-cluster"
	clientID  = "test-publisher"
	channel   = "test"
)

func main() {
	js, err := os.Open("data.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer js.Close()

	natsConnection, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Println("", err)
	}

	jsByte, _ := ioutil.ReadAll(js)

	var data []domain.OrderEntity
	err = json.Unmarshal(jsByte, &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, value := range data {
		order, _ := json.Marshal(value)

		err = natsConnection.Publish(channel, order)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = natsConnection.Close()
	if err != nil {
		log.Println(err)
	}
}
