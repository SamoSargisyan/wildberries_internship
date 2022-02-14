package main

import (
	"l0/internal/service"
	"log"
)

func main() {
	err := service.Bootstrap()
	if err != nil {
		log.Fatalln(err)
	}
}
