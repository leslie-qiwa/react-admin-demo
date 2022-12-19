package models

import (
	"time"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TableName Database Table Name of this model
func (cat *Category) TableName() string {
	return "categories"
}

type Customer struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	Zipcode    string    `json:"zipcode"`
	City       string    `json:"city"`
	StateAbbr  string    `json:"stateAbbr"`
	Avatar     string    `json:"avatar"`
	Birthday   time.Time `json:"birthday"`
	FirstSeen  time.Time `json:"first_seen"`
	LastSeen   time.Time `json:"last_seen"`
	HasOrdered bool      `json:"has_ordered"`
	//LatestPurchase string `json:"latest_purchase"`
	HasNewsletter bool `json:"has_newsletter"`
	//groups: array
	NbCommands int `json:"nb_commands"`
	TotalSpent int `json:"total_spent"`
}

type Product struct {
	ID          int     `json:"id"`
	CategoryID  int     `json:"category_id"`
	Reference   string  `json:"reference"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	Price       float32 `json:"price"`
	Thumbnail   string  `json:"thumbnail"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

type Invoice struct {
	ID           int       `json:"id"`
	Date         time.Time `json:"date"`
	CommandID    int       `json:"command_id"`
	CustomerID   int       `json:"customer_id"`
	TotalExTaxes float32   `json:"total_ex_taxes"`
	DeliveryFees float32   `json:"delivery_fees"`
	TaxRate      float32   `json:"tax_rate"`
	Taxes        float32   `json:"taxes"`
	Total        float32   `json:"total"`
}

type Review struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	Status     Status    `json:"status"`
	CommandID  int       `json:"command_id"`
	ProductID  int       `json:"product_id"`
	CustomerID int       `json:"customer_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
}
