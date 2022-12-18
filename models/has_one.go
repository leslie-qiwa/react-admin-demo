package models

// User has one CreditCardOne, CreditCardID is the foreign key

type UserOne struct {
	CreditCard CreditCard
}

type CreditCardOne struct {
	Number string
	UserID uint
}

type ProductOne struct {
	Category Category
}

type CategoryOne struct {
	Name string
}
