package dto

type PaymentRequest struct {
	CustomerPhoneNo string `json:"customer_phone_no"`
	Total           int    `json:"total"`
}
