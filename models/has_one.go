package models

// User has one CreditCardOne, CreditCardID is the foreign key

type ProductOne struct {
	Category Category
}

type CategoryOne struct {
	Name string
}
