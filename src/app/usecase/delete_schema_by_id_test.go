package usecase

import (
	"testing"

	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func TestDeleteSchemaByID(ts *testing.T) {
	ts.Run("it should be able to filter the cluster schemas", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		deleteSchemaByID := NewDeleteSchemaByID(repo)
		createCluster := NewCreateCluster(repo)
		createSchema := NewCreateSchema(repo)

		c, _ := createCluster.Exec(name)
		createSchema.Exec(name, map[string]string{"name": "Shinji", "age": "14"})
		s, _ := createSchema.Exec(name, map[string]string{"name": "Asuka", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Rei", "age": "14"})

		err := deleteSchemaByID.Exec(name, s.ID)
		if err != nil || c.Schemas[1].ID == s.ID {
			t.Error(err)
		}
	})
}
