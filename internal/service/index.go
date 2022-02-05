package service

import (
	"github.com/jmoiron/sqlx"
	"l0/config"
	"l0/internal/domain"
	"l0/internal/repository/order_repository"
	"l0/pkg/db/postgres"
	"log"
)

func Bootstrap(cfg *config.Config, orderUID string) (*domain.OrderEntity, error) {

	db, err := postgres.NewConnection(cfg.Database.Combine)
	if err != nil {
		return nil, err
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	st := order_repository.Init(db)

	order, err := st.Order().FindByID(orderUID)

	if err != nil {
		return nil, err
	}

	log.Print(order.DateCreated, "success")
	return order, err
}
