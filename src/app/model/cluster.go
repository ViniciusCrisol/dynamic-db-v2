package model

import (
	"time"

	"github.com/google/uuid"
)

type Cluster struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Schemas   []*Schema `json:"schemas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewCluster returns a cluster using its name.
func NewCluster(name string) *Cluster {
	now := time.Now()
	return &Cluster{
		ID:        uuid.NewString(),
		Name:      name,
		Schemas:   []*Schema{},
		CreatedAt: now,
		UpdatedAt: now,
	}
}
