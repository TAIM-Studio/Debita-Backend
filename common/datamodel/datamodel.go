package datamodel

import (
	"time"

	"github.com/google/uuid"
)

type Abo struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Version uint      `json:"version"`
	Price   float64   `json:"price"`
}

type Debtor struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Abos      []Abo     `json:"abos" gorm:"many2many:debtor_abos;"`
	Debts     []Debt    `json:"debts" gorm:"many2many:debtor_debts;"`
}

type Debt struct {
	Id     uuid.UUID `json:"id"`
	Reason string    `json:"reason"`
	Amount float64   `json:"amount"`
}

type Payement struct {
	Id     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}
