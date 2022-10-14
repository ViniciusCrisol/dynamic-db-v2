package repo

import (
	"errors"
	"fmt"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type clusterMem struct {
	baseURL  string
	clusters []*model.Cluster
}

func NewClusterMem() app.ClusterRepo {
	return &clusterMem{
		baseURL:  ":mem/clusters",
		clusters: []*model.Cluster{},
	}
}

func (repo *clusterMem) AssembleURL(name string) string {
	return fmt.Sprintf("%s/%s", repo.baseURL, name)
}

func (repo *clusterMem) Create(cluster *model.Cluster) error {
	for _, c := range repo.clusters {
		if c.URL == cluster.URL {
			return errors.New("unique constraint violation")
		}
	}
	repo.clusters = append(repo.clusters, cluster)
	return nil
}

func (repo *clusterMem) Update(cluster *model.Cluster) error {
	for i, c := range repo.clusters {
		if c.URL == cluster.URL {
			repo.clusters[i] = cluster
			return nil
		}
	}
	return errors.New("invalid cluster")
}

func (repo *clusterMem) Find(url string) (*model.Cluster, error) {
	for _, c := range repo.clusters {
		if c.URL == url {
			return c, nil
		}
	}
	return nil, nil
}
