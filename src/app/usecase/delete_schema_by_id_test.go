package usecase

import (
	"testing"

	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func TestDeleteSchemaByID(ts *testing.T) {
	ts.Run("it should be able to delete a schema", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		deleteSchemaByID := NewDeleteSchemaByID(repo)
		createCluster := NewCreateCluster(repo)
		createSchema := NewCreateSchema(repo)

		c, _ := createCluster.Exec(name)
		s, _ := createSchema.Exec(name, map[string]string{"name": "Asuka", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Shinji", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Rei", "age": "14"})

		err := deleteSchemaByID.Exec(name, s.ID)
		if err != nil || c.Schemas[0].ID == s.ID {
			t.Error(err)
		}
	})

	ts.Run("it should not be able to delete a schema of a non-existent cluster", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		usecase := NewDeleteSchemaByID(repo)

		err := usecase.Exec(name, "generic-ID")
		if err == nil {
			t.Error()
		}
	})

	ts.Run("it should not be able to delete a non-existent schema", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		deleteSchemaByID := NewDeleteSchemaByID(repo)
		createCluster := NewCreateCluster(repo)
		createCluster.Exec(name)

		err := deleteSchemaByID.Exec(name, "generic-ID")
		if err == nil {
			t.Error()
		}
	})
}
