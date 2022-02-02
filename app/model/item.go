package model

type ItemEntity struct {
	ChrtId      uint64 `db:"chrtId"`
	TrackNumber string `db:"trackNumber"`
	Price       int64  `db:"price"`
	Rid         string `db:"rid"`
	Name        string `db:"name"`
	Sale        int64  `db:"sale"`
	Size        string `db:"size"`
	TotalPrice  int64  `db:"totalPrice"`
	NmID        uint64 `db:"nmID"`
	Status      int    `db:"status"`
	OrderUID    string `db:"orderUID"`
}
