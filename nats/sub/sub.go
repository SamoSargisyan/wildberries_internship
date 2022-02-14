package natssub

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	//"github.com/patrickmn/go-cache"
	"l0/config"
	"l0/internal/cache"
	"l0/internal/domain"
	"l0/internal/repository/order_repository"
	"log"
	"path/filepath"
	"time"
)

func HandleNewOrder(order *domain.OrderEntity, store *order_repository.OrderStore, cacheLocal *cache.LocalCache) error {
	err := store.Order().InsertOrder(order)
	if err != nil {
		log.Println("In HandleNewOrder: ", err)
		return err
	}

	cacheLocal.Set(order.OrderUID, order, 5*time.Minute)

	return nil
}

func Sub(conn stan.Conn, store *order_repository.OrderStore, cacheLocal *cache.LocalCache) stan.Subscription {
	mainConfigFile, _ := filepath.Abs("./config/services.yml")
	cfg := config.GetNatsConfigurations(mainConfigFile)

	handler := func(msg *stan.Msg) {
		var data domain.OrderEntity

		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Printf("error while decoding data from nats-pub: %v ", err)

			err := msg.Ack()
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		validate := validator.New()
		err = validate.Struct(data)
		if err != nil {

			err := msg.Ack()
			if err != nil {
				log.Println(err)
				return
			}

			return
		}

		err = HandleNewOrder(&data, store, cacheLocal)
		if err != nil {
			log.Printf("error while inserting(to db) data from nats-pub: %v ", err)

			err := msg.Ack()
			if err != nil {
				log.Println(err)
				return
			}

			return
		}

		err = msg.Ack()
		if err != nil {
			log.Printf("failed ACK msg: %d", msg.Sequence)
			return
		}
	}

	sub, err := conn.Subscribe(
		cfg.Nats.Channel,
		handler,
		stan.DurableName("durable-name"),
		stan.SetManualAckMode(),
	)
	if err != nil {
		log.Fatal(err)
	}

	return sub
}
