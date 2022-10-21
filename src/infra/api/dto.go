package api

type DefaultResponse struct {
	Data    any    `json:"data,omitempty"`
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type CreateClusterDTO struct {
	Name string `json:"name"`
}

type CreateSchemaDTO struct {
	Name    string            `json:"cluster"`
	Content map[string]string `json:"content"`
}
