package usecase

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
)

type DeleteSchemaByID struct {
	repo app.ClusterRepo
}

// NewDeleteSchemaByID returns a new delete schema usecase.
func NewDeleteSchemaByID(repo app.ClusterRepo) *DeleteSchemaByID {
	return &DeleteSchemaByID{repo}
}

// Exec gets the cluster by its name and deletes its schemas using the schema ID.
func (ucs *DeleteSchemaByID) Exec(name string, id string) error {
	c, err := ucs.repo.Find(name)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("INVALID-CLUSTER")
	}
	for i, schema := range c.Schemas {
		if schema.ID == id {
			c.Schemas = append(c.Schemas[:i], c.Schemas[i+1:]...)
			return ucs.repo.Update(c)
		}
	}
	return errors.New("INVALID-SCHEMA")
}
