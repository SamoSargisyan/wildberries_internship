package order_repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type OrderStore struct {
	orderRepository *OrderRepository
	db              *sqlx.DB
}

func Init(db *sqlx.DB) *OrderStore {
	return &OrderStore{
		db: db,
	}
}

func (s *OrderStore) Order() *OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		s,
	}

	return s.orderRepository
}
