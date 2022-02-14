package order_repository

import (
	"errors"
	"l0/internal/domain"
	"log"
)

type OrderRepository struct {
	*OrderStore
}

func (r *OrderRepository) FindAll() []domain.OrderEntity {
	var orders []domain.OrderEntity
	query := `SELECT *
	   		  FROM orders
			  ORDER BY date_created LIMIT 10000;`

	err := r.db.Select(&orders, query)
	if err != nil {
		log.Fatal(err)
	}

	return orders
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

func (r *OrderRepository) InsertOrder(domain *domain.OrderEntity) error {

	orderID, err := r.checkUnique(domain.OrderUID)
	if err != nil {
		return err
	}

	if orderID == domain.OrderUID {
		return errors.New("not unique OrderUID")
	}

	err = r.insertOrder(domain)
	if err != nil {
		return err
	}

	err = r.insertDelivery(&domain.Delivery, domain.OrderUID)
	if err != nil {
		return err
	}

	err = r.insertPayment(&domain.Payment, domain.OrderUID)
	if err != nil {
		return err
	}

	err = r.insertItems(domain.Items, domain.OrderUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) checkUnique(orderUID string) (string, error) {
	var id string
	query := `SELECT order_uid FROM orders where order_uid = $1`

	err := r.db.Get(&id, query, orderUID)
	if err != nil {
		return id, nil
	}

	return id, nil
}

func (r *OrderRepository) insertOrder(order *domain.OrderEntity) error {
	sqlStatement := `INSERT INTO orders
				(
					order_uid,
					track_number,
					entry,
					locale,
					internal_signature,
					customer_id,
					delivery_service,
					shardkey,
					sm_id,
					date_created,
					oof_shard
				)
			    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(sqlStatement, order.OrderUID,
		order.TrackNumber, order.Entry, order.Local, order.InternalSignature,
		order.CustomerID, order.DeliveryService,
		order.Shardkey, order.SmID, order.DateCreated, order.OofShard)

	if err != nil {
		panic(err)
	}

	return nil
}

func (r *OrderRepository) insertDelivery(delivery *domain.DeliveryEntity, orderUID string) error {
	sqlStatement := `INSERT INTO deliveries
		(
			name,
			phone,
			zip,
			city,
			address,
			region,
			email,
			order_uid
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.Exec(sqlStatement, delivery.Name, delivery.Phone, delivery.Zip, delivery.City,
		delivery.Address, delivery.Region, delivery.Email, orderUID)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) insertItems(items []domain.ItemEntity, orderUID string) error {
	sqlStatement := `INSERT INTO items
		(
			chrt_id,
			track_number,
			price,
			rid,
			name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status,
			order_uid
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	for _, item := range items {
		_, err := r.db.Exec(sqlStatement, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale,
			item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status, orderUID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) insertPayment(payment *domain.PaymentEntity, orderUID string) error {
	sqlStatement := `INSERT INTO payments
		(
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee,
			order_uid
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(sqlStatement, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider,
		payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal,
		payment.CustomFee, orderUID)

	if err != nil {
		return err
	}

	return nil
}
