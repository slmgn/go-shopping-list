package models

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
