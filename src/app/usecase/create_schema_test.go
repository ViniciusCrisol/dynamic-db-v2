package usecase

import (
	"testing"

	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func TestCreateSchemas(ts *testing.T) {
	ts.Run("it should be able to filter the cluster schemas", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		createSchema := NewCreateSchema(repo)
		createCluster := NewCreateCluster(repo)
		createCluster.Exec(name)

		s, err := createSchema.Exec(name, map[string]string{"name": "Asuka", "age": "14"})
		if err != nil ||
			s.Content["age"] != "14" ||
			s.Content["name"] != "Asuka" {
			t.Error(err)
		}
	})

	ts.Run("it should not be able to create a schema of a non-existent cluster", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		usecase := NewCreateSchema(repo)

		_, err := usecase.Exec(name, map[string]string{"name": "Asuka", "age": "14"})
		if err == nil {
			t.Error()
		}
	})
}
