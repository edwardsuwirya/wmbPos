package dto

type CustomerOrderRequest struct {
	CustomerName string
	TableId      string
	Orders       []CustomerOrderDetailRequest
}

type CustomerOrderDetailRequest struct {
	MenuId string
	Qty    int
}
