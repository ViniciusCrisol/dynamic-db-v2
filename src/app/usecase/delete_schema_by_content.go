package usecase

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type DeleteSchemaByContent struct {
	repo app.ClusterRepo
}

// NewDeleteSchemaByContent returns a new delete schema usecase.
func NewDeleteSchemaByContent(repo app.ClusterRepo) *DeleteSchemaByContent {
	return &DeleteSchemaByContent{repo}
}

// Exec gets the cluster by its name and deletes its schemas using the content.
func (ucs *DeleteSchemaByContent) Exec(name string, content map[string]string) error {
	c, err := ucs.repo.Find(name)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("INVALID-CLUSTER")
	}
	s := []*model.Schema{}
	for _, schema := range c.Schemas {
		keep := true
		for k, v := range content {
			if schema.Content[k] == v {
				keep = false
			}
		}
		if keep {
			s = append(s, schema)
		}
	}
	c.Schemas = s
	return ucs.repo.Update(c)
}
