package order_repository

import (
	"l0/internal/domain"
)

type OrderRepository struct {
	*OrderStore
}

func (r *OrderRepository) FindByID(orderUID string) (*domain.OrderEntity, error) {
	order, err := r.getOrder(orderUID)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *OrderRepository) getOrder(orderUID string) (*domain.OrderEntity, error) {
	order := domain.OrderEntity{}
	query := `SELECT * FROM orders where order_uid = $1`

	err := r.db.Get(&order, query, orderUID)
	if err != nil {
		return &order, err
	}

	return &order, nil
}
