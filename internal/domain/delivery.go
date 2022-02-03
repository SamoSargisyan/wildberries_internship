package domain

type DeliveryEntity struct {
	ID       uint64 `db:"ID"`
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	Zip      string `db:"zip"`
	City     string `db:"city"`
	Address  string `db:"address"`
	Region   string `db:"region"`
	Email    string `db:"email"`
	OrderUID string `db:"order_uid"`
}
