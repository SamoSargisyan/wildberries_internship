package domain

type PaymentEntity struct {
	Transaction  string `db:"transaction" json:"transaction" validate:"required"`
	RequestID    string `db:"request_id" json:"request_id" validate:"required"`
	Currency     string `db:"currency" json:"currency" validate:"required,max=128"`
	Provider     string `db:"provider" json:"provider" validate:"required,max=128"`
	Amount       uint64 `db:"amount" json:"amount" validate:"required"`
	PaymentDt    uint64 `db:"payment_dt" json:"payment_dt" validate:"required"`
	Bank         string `db:"bank" json:"bank" validate:"required,max=128"`
	DeliveryCost uint64 `db:"delivery_cost" json:"delivery_cost" validate:"omitempty"`
	GoodsTotal   uint64 `db:"goods_total" json:"goods_total" validate:"required"`
	CustomFee    uint64 `db:"custom_fee" json:"custom_fee" validate:"omitempty"`
}
