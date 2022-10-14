package model

import (
	"time"

	"github.com/google/uuid"
)

type Schema struct {
	ID        string
	Content   map[string]string
	CreatedAt time.Time
	UpdatedAt time.Time
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
