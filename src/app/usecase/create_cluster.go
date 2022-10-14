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
	url := ucs.repo.AssembleURL(name)
	clusterWithSameURL, err := ucs.repo.Find(url)
	if err != nil {
		return nil, err
	}
	if clusterWithSameURL != nil {
		return nil, errors.New("CLUSTER-URL-IN-USE")
	}
	c := model.NewCluster(url)
	err = ucs.repo.Create(c)
	return c, err
}
