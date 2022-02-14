package domain

type ItemEntity struct {
	ChrtId      uint64 `db:"chrt_id" json:"chrt_id" validate:"required"`
	TrackNumber string `db:"track_number" json:"track_number" validate:"required,max=256"`
	Price       int64  `db:"price" json:"price" validate:"required"`
	Rid         string `db:"rid" json:"rid" validate:"required"`
	Name        string `db:"name" json:"name" validate:"required,max=128"`
	Sale        int64  `db:"sale" json:"sale" validate:"required"`
	Size        string `db:"size" json:"size" validate:"required"`
	TotalPrice  int64  `db:"total_price" json:"total_price" validate:"required"`
	NmID        uint64 `db:"nm_id" json:"nm_id" validate:"required"`
	Brand       string `db:"brand" json:"brand" validate:"required,max=256"`
	Status      int    `db:"status" json:"status" validate:"required"`
}
