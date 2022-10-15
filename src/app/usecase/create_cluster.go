package usecase

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type CreateCluster struct {
	repo app.ClusterRepo
}

// NewCreateCluster returns a new create cluster usecase.
func NewCreateCluster(repo app.ClusterRepo) *CreateCluster {
	return &CreateCluster{repo}
}

// Exec creates a new cluster. It checks if the name is available and returns an error if
// it's not.
func (ucs *CreateCluster) Exec(name string) (*model.Cluster, error) {
	clusterWithSameName, err := ucs.repo.Find(name)
	if err != nil {
		return nil, err
	}
	if clusterWithSameName != nil {
		return nil, errors.New("CLUSTER-NAME-IN-USE")
	}
	c := model.NewCluster(name)
	return c, ucs.repo.Create(c)
}
