package repo

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type cluster struct {
	baseURL          string
	clusterSemaphore map[string]*sync.Mutex
}

func NewCluster(baseURL string) app.ClusterRepo {
	return &cluster{
		baseURL:          baseURL,
		clusterSemaphore: map[string]*sync.Mutex{},
	}
}

func (repo *cluster) AssembleURL(name string) string {
	const F_EXTENSION = "gob"
	return fmt.Sprintf("%s/%s.%s", repo.baseURL, name, F_EXTENSION)
}

func (repo *cluster) Create(cluster *model.Cluster) error {
	return repo.overwriteClusterFile(cluster)
}

func (repo *cluster) Update(cluster *model.Cluster) error {
	return repo.overwriteClusterFile(cluster)
}

func (repo *cluster) Find(url string) (*model.Cluster, error) {
	repo.lockClusterFile(url)
	defer repo.unlockClusterFile(url)

	_, status := os.Stat(url)
	fExists := os.IsNotExist(status)
	if fExists {
		return nil, nil
	}

	f, err := os.Open(url)
	if err != nil {
		// TODO: Handle it.
		return nil, err
	}
	defer f.Close()

	cluster := &model.Cluster{}
	decoder := gob.NewDecoder(f)
	err = decoder.Decode(cluster)
	if err != nil {
		// TODO: Handle it.
		return nil, err
	}
	return cluster, nil
}

func (repo *cluster) lockClusterFile(url string) {
	smr, ok := repo.clusterSemaphore[url]
	if !ok {
		smr = &sync.Mutex{}
		repo.clusterSemaphore[url] = smr
	}
	smr.Lock()
}

func (repo *cluster) unlockClusterFile(url string) {
	smr := repo.clusterSemaphore[url]
	smr.Unlock()
}

func (repo *cluster) overwriteClusterFile(cluster *model.Cluster) error {
	repo.lockClusterFile(cluster.URL)
	defer repo.unlockClusterFile(cluster.URL)

	f, err := os.Create(cluster.URL)
	if err != nil {
		// TODO: Handle it.
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(cluster)
	if err != nil {
		// TODO: Handle it.
	}
	return err
}
