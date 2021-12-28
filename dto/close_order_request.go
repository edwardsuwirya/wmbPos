package dto

type CloseOrderRequest struct {
	BillNo        string
	PaymentMethod string
	PhoneNo       string
	Total         int
}
