package dto

type OpoHttpResponse struct {
	Message string
	Data    OpoReceipt
}

type OpoReceipt struct {
	ReceiptId string `json:"receipt_id"`
}
