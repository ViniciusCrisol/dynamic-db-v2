package usecase

import (
	"testing"

	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func TestDeleteSchemaByValue(ts *testing.T) {
	ts.Run("it should be able to delete a schema", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		deleteSchemaByValue := NewDeleteSchemaByContent(repo)
		createCluster := NewCreateCluster(repo)
		createSchema := NewCreateSchema(repo)

		c, _ := createCluster.Exec(name)
		s, _ := createSchema.Exec(name, map[string]string{"name": "Asuka", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Shinji", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Rei", "age": "14"})

		err := deleteSchemaByValue.Exec(name, map[string]string{"name": "Asuka"})
		if err != nil || c.Schemas[0].ID == s.ID {
			t.Error(err)
		}
	})

	ts.Run("it should be able to many schemas", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		deleteSchemaByValue := NewDeleteSchemaByContent(repo)
		createCluster := NewCreateCluster(repo)
		createSchema := NewCreateSchema(repo)

		c, _ := createCluster.Exec(name)
		createSchema.Exec(name, map[string]string{"name": "Shinji", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Asuka", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Rei", "age": "14"})

		err := deleteSchemaByValue.Exec(name, map[string]string{"age": "14"})
		if err != nil || len(c.Schemas) != 0 {
			t.Error(err)
		}
	})

	ts.Run("it should not be able to delete a schema of a non-existent cluster", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		usecase := NewDeleteSchemaByContent(repo)

		err := usecase.Exec(name, map[string]string{"name": "Asuka"})
		if err == nil {
			t.Error()
		}
	})
}
