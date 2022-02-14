package service

import (
	"l0/http_server"
	"l0/internal/cache"
	"l0/internal/repository/order_repository"
	"l0/nats/sub"
)

func Bootstrap() error {
	db := getDb()
	defer db.Close()

	store := order_repository.Init(db)
	orders := store.Order().FindAll()

	cacheLocal := cache.NewCache(orders)

	natsConnection := getStan()
	defer natsConnection.Close()

	subscription := natssub.Sub(natsConnection, store, cacheLocal)
	defer subscription.Unsubscribe()

	server := http_server.InitServer(cacheLocal)
	return server.Start()
}
