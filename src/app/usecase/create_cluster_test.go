package usecase

import (
	"testing"

	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func TestActivateUserUsecase(ts *testing.T) {
	ts.Run("it should be able to create a cluster", func(t *testing.T) {
		repo := repo.NewClusterMem()
		usecase := NewCreateCluster(repo)

		_, err := usecase.Exec("cluster-name")
		if err != nil {
			t.Error(err)
		}
	})

	ts.Run("it should not be able to create a cluster with a duplicated name", func(t *testing.T) {
		name := "cluster-name"
		repo := repo.NewClusterMem()
		usecase := NewCreateCluster(repo)

		usecase.Exec(name)
		_, err := usecase.Exec(name)
		if err == nil {
			t.Error()
		}
	})
}
