package usecase

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type CreateSchema struct {
	repo app.ClusterRepo
}

// CreateSchema returns a new create schema usecase.
func NewCreateSchema(repo app.ClusterRepo) *CreateSchema {
	return &CreateSchema{repo}
}

// Exec creates a new schema into the cluster and stores it.
func (ucs *CreateSchema) Exec(name string, content map[string]string) (*model.Schema, error) {
	c, err := ucs.repo.Find(name)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("INVALID-CLUSTER")
	}
	s := model.NewSchema(content)
	c.Schemas = append(c.Schemas, s)
	return s, ucs.repo.Update(c)
}
