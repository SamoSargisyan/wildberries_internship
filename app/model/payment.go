package model

type PaymentEntity struct {
	Transaction  string `db:"transaction"`
	RequestID    string `db:"requestID"`
	Currency     string `db:"currency"`
	Provider     string `db:"provider"`
	Amount       int64  `db:"amount"`
	PaymentDt    int64  `db:"paymnetDt"`
	Bank         string `db:"bank"`
	DeliveryCost int64  `db:"deliveryCost"`
	GoodsTotal   int64  `db:"goodsTotal"`
	CustomFee    int64  `db:"customFee"`
	OrderUID     string `db:"orderUID"`
}
