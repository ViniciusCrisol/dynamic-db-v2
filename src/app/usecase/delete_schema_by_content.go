package usecase

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type DeleteSchemaByValue struct {
	repo app.ClusterRepo
}

// NewDeleteSchemaByValue returns a new delete schema usecase.
func NewDeleteSchemaByValue(repo app.ClusterRepo) *DeleteSchemaByValue {
	return &DeleteSchemaByValue{repo}
}

// Exec gets the cluster by its name and deletes its schemas using the content key and
// value.
func (ucs *DeleteSchemaByValue) Exec(name string, content map[string]string) error {
	c, err := ucs.repo.Find(name)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("INVALID-CLUSTER")
	}
	s := []*model.Schema{}
	for _, schema := range c.Schemas {
		keepSchema := true
		for k, v := range content {
			if schema.Content[k] == v {
				keepSchema = false
			}
		}
		if keepSchema {
			s = append(s, schema)
		}
	}
	c.Schemas = s
	return ucs.repo.Update(c)
}
