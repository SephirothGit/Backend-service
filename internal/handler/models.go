package handler

type UpdateOrderStatus struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
