package models

import (
	"gorm.io/gorm"
	"time"
)

// User has many CreditCards, UserID is the foreign key
type User struct {
	ID          uint         `json:"id" gorm:"primary_key"`
	Name        string       `json:"name"`
	CreditCards []CreditCard `json:"credit_cards"`
}

type CreditCard struct {
	gorm.Model

	ID     uint   `json:"id" gorm:"primary_key"`
	Number string `json:"number"`
	UserID uint   `json:"user_id"`
}

type Status string

const (
	StatusOrdered   Status = "ordered"
	StatusDelivered Status = "delivered"
	StatusCanceled  Status = "canceled"
	StatusAccepted  Status = "accepted"
	StatusRejected  Status = "rejected"
	StatusPending   Status = "pending"
)

type Basket struct {
	gorm.Model `json:"-"`

	ID        uint `json:"-" gorm:"primary_key"`
	ProductID int  `json:"product_id"`
	Quantity  int  `json:"quantity"`
	CommandID uint `json:"-" gorm:"ForeignKey:ID"`
}

// Command has many Basket, Command ID is the foreign key
type Command struct {
	ID           int       `json:"id" gorm:"primary_key"`
	Reference    string    `json:"reference"`
	Date         time.Time `json:"date"`
	CustomerID   int       `json:"customer_id"`
	Baskets      []Basket  `json:"basket"`
	TotalExTaxes float32   `json:"total_ex_taxes"`
	DeliveryFees float32   `json:"delivery_fees"`
	TaxRate      float32   `json:"tax_rate"`
	Taxes        float32   `json:"taxes"`
	Total        float32   `json:"total"`
	Status       Status    `json:"status"`
	Returned     bool      `json:"returned"`
}
