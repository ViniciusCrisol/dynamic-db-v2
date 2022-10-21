package usecase

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type FilterSchemas struct {
	repo app.ClusterRepo
}

// NewFilterSchemas returns a new filter cluster schema usecase.
func NewFilterSchemas(repo app.ClusterRepo) *FilterSchemas {
	return &FilterSchemas{repo}
}

// Exec gets the cluster by its name and filters it using the content key and value.
func (ucs *FilterSchemas) Exec(name string, content map[string]string) ([]*model.Schema, error) {
	c, err := ucs.repo.Find(name)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("INVALID-CLUSTER")
	}
	s := []*model.Schema{}
	for _, schema := range c.Schemas {
		keepSchema := true
		for k, v := range content {
			if schema.Content[k] != v {
				keepSchema = false
			}
		}
		if keepSchema {
			s = append(s, schema)
		}
	}
	return s, nil
}
