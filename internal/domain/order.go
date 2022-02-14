package domain

import (
	"time"
)

type OrderEntity struct {
	OrderUID          string    `db:"order_uid" json:"order_uid" validate:"required"`
	TrackNumber       string    `db:"track_number" json:"track_number" validate:"required,max=128"`
	Entry             string    `db:"entry" json:"entry" json:"entry" validate:"required,max=128"`
	Local             string    `db:"locale" json:"locale" validate:"required,max=128"`
	InternalSignature string    `db:"internal_signature" json:"internal_signature" validate:"omitempty,max=128"`
	CustomerID        string    `db:"customer_id" json:"customer_id" validate:"required,max=128"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service" validate:"required,max=128"`
	Shardkey          string    `db:"shardkey" json:"shardkey" validate:"required,max=128"`
	SmID              uint64    `db:"sm_id" json:"sm_id" validate:"required"`
	DateCreated       time.Time `db:"date_created" json:"date_created" validate:"required"`
	OofShard          string    `db:"oof_shard" json:"oof_shard" validate:"required,max=128"`

	Delivery DeliveryEntity `db:"delivery" json:"delivery"`
	Payment  PaymentEntity  `db:"payment" json:"payment"`
	Items    []ItemEntity   `db:"items" validate:"required,dive,required" json:"items"`
}
