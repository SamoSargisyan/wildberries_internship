package domain

import (
	"time"
)

type OrderEntity struct {
	OrderUID          string    `db:"order_uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	local             string    `db:"local"`
	InternalSignature string    `db:"internal_signature"`
	CustomerID        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	Shardkey          string    `db:"shardkey"`
	SmID              uint64    `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}
