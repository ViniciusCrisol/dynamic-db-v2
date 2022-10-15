package repo

import (
	"errors"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type clusterMem struct {
	basePath string
	clusters []*model.Cluster
}

func NewClusterMem() app.ClusterRepo {
	return &clusterMem{
		basePath: ":mem/clusters",
		clusters: []*model.Cluster{},
	}
}

func (repo *clusterMem) Create(cluster *model.Cluster) error {
	for _, c := range repo.clusters {
		if c.Name == cluster.Name {
			return errors.New("CLUSTER-NAME-IN-USE")
		}
	}
	repo.clusters = append(repo.clusters, cluster)
	return nil
}

func (repo *clusterMem) Update(cluster *model.Cluster) error {
	for i, c := range repo.clusters {
		if c.Name == cluster.Name {
			repo.clusters[i] = cluster
			return nil
		}
	}
	return errors.New("INVALID-CLUSTER")
}

func (repo *clusterMem) Find(name string) (*model.Cluster, error) {
	for _, c := range repo.clusters {
		if c.Name == name {
			return c, nil
		}
	}
	return nil, nil
}
