package model

import (
	"time"

	"github.com/google/uuid"
)

type Cluster struct {
	ID        string
	URL       string
	schemas   []*Schema
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewCluster returns a new cluster.
func NewCluster(url string) *Cluster {
	now := time.Now()
	return &Cluster{
		ID:        uuid.NewString(),
		URL:       url,
		schemas:   []*Schema{},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// AppendSchema stores a schema in the cluster.
func (c *Cluster) AppendSchema(schema *Schema) {
	c.schemas = append(c.schemas, schema)
}
