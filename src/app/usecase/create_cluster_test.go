package usecase

import (
	"testing"

	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func TestCreateCluster(ts *testing.T) {
	ts.Run("it should be able to create a cluster", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		create := NewCreateCluster(repo)

		_, err := create.Exec(name)
		if err != nil {
			t.Error(err)
		}
	})

	ts.Run("it should not be able to create a cluster with a duplicated name", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		create := NewCreateCluster(repo)
		create.Exec(name)

		_, err := create.Exec(name)
		if err == nil {
			t.Error()
		}
	})
}
