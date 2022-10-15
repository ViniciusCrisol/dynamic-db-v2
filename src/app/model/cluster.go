package model

import (
	"time"

	"github.com/google/uuid"
)

type Cluster struct {
	ID        string
	Name      string
	Schemas   []*Schema
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewCluster returns a new cluster.
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

// AppendSchema stores a schema in the cluster.
func (c *Cluster) AppendSchema(schema *Schema) {
	c.Schemas = append(c.Schemas, schema)
}
