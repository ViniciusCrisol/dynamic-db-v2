package usecase

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type CreateCluster struct {
	repo app.ClusterRepo
}

func NewCreateCluster(repo app.ClusterRepo) *CreateCluster {
	return &CreateCluster{repo}
}

func (ucs *CreateCluster) Exec(name string) (*model.Cluster, error) {
	clusterWithSameName, err := ucs.repo.Find(name)
	if err != nil {
		return nil, err
	}
	if clusterWithSameName != nil {
		return nil, errors.New("CLUSTER-NAME-IN-USE")
	}
	c := model.NewCluster(name)
	err = ucs.repo.Create(c)
	return c, err
}
