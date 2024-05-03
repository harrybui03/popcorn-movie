package model

import "github.com/google/uuid"

type CreateTicket struct {
	SeatID        uuid.UUID
	TransactionID uuid.UUID
	ShowTimeID    uuid.UUID
	Price         float64
	IsBooked      bool
}

type CreateFoodOrderLine struct {
	FoodID        uuid.UUID
	Quantity      int
	TransactionID uuid.UUID
}
