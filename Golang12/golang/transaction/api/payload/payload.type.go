package payload

type OrderPayload struct {
	CustomerID int32 `json:"customer_id"`
	ProductID  int32 `json:"product_id"`
	Quantity   int32 `json:"quantity"`
}
