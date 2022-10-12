package api

type DefaultResponse struct {
	Data    any    `json:"data,omitempty"`
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
