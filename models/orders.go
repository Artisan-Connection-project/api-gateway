package models

type UpdateOrderStatus struct {
	Status string `json:"status"`
}

type PayOrderRequest struct {
	PaymentMethod string `json:"payment_method"`
	CardNumber    string `json:"card_number"`
	ExpiryDate    string `json:"expiry_date"`
	Cvv           string `json:"cvv"`
}

type UpdateShippingInfoRequest struct {
	TrackingNumber        string `json:"tracking_number"`
	Carrier               string `json:"carrier"`
	EstimatedDeliveryDate string `json:"estimated_delivery_date"`
}
