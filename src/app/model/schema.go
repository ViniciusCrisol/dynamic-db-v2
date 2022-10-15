package model

import (
	"time"

	"github.com/google/uuid"
)

type Schema struct {
	ID        string            `json:"id"`
	Content   map[string]string `json:"content"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// NewCluster returns a new schema. It could be stored on a cluster.
func NewSchema(content map[string]string) *Schema {
	now := time.Now()
	return &Schema{
		ID:        uuid.NewString(),
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
