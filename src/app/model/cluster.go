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

// FilterSchemas filters the cluster schemas using the content key and value.
func (c *Cluster) FilterSchemas(k, v string) []*Schema {
	schemas := []*Schema{}
	for _, schema := range c.Schemas {
		if schema.Content[k] == v {
			schemas = append(schemas, schema)
		}
	}
	return schemas
}
