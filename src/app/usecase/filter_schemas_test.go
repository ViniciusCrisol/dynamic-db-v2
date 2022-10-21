package usecase

import (
	"testing"

	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func TestFilterSchemas(ts *testing.T) {
	ts.Run("it should be able to filter the schemas", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		filter := NewFilterSchemas(repo)
		createSchema := NewCreateSchema(repo)
		createCluster := NewCreateCluster(repo)

		createCluster.Exec(name)
		createSchema.Exec(name, map[string]string{"name": "Shinji", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Asuka", "age": "14"})
		createSchema.Exec(name, map[string]string{"name": "Rei", "age": "14"})

		s, err := filter.Exec(name, "name", "Asuka")
		if err != nil ||
			s[0].Content["age"] != "14" ||
			s[0].Content["name"] != "Asuka" {
			t.Error(err)
		}
	})

	ts.Run("it should not be able to filter the schemas of a non-existent cluster", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		usecase := NewFilterSchemas(repo)

		_, err := usecase.Exec(name, "name", "Asuka")
		if err == nil {
			t.Error()
		}
	})
}
