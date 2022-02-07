package domain

type DeliveryEntity struct {
	Name    string `db:"name" json:"name" validate:"required,max=128"`
	Phone   string `db:"phone" json:"phone" validate:"required,max=16" faker:"len=16"`
	Zip     string `db:"zip" json:"zip" validate:"required,max=128"`
	City    string `db:"city" json:"city" validate:"required,max=128"`
	Address string `db:"address" json:"address" validate:"required,max=256"`
	Region  string `db:"region" json:"region" validate:"required,max=256"`
	Email   string `db:"email" json:"email" validate:"required,max=128"`
}
