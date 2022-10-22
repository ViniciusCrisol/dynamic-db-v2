package repo

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/model"
)

type cluster struct {
	basePath         string
	clusterSemaphore map[string]*sync.Mutex
}

func NewCluster(basePath string) app.ClusterRepo {
	return &cluster{
		basePath:         basePath,
		clusterSemaphore: map[string]*sync.Mutex{},
	}
}

func (repo *cluster) Create(cluster *model.Cluster) error {
	path := repo.assemblePath(cluster.Name)
	f, err := repo.findByPath(path)
	if err != nil {
		return err
	}
	if f != nil {
		return errors.New("CLUSTER-NAME-IN-USE")
	}
	return repo.overwriteClusterFile(path, cluster)
}

func (repo *cluster) Update(cluster *model.Cluster) error {
	path := repo.assemblePath(cluster.Name)
	f, err := repo.findByPath(path)
	if err != nil {
		return err
	}
	if f == nil {
		return errors.New("INVALID-CLUSTER")
	}
	return repo.overwriteClusterFile(path, cluster)
}

func (repo *cluster) Find(name string) (*model.Cluster, error) {
	path := repo.assemblePath(name)
	return repo.findByPath(path)
}

// findByPath searches for a cluster by its path. If it exists, it will be returned.
// Otherwise, a nil pointer will be returned.
func (repo *cluster) findByPath(path string) (*model.Cluster, error) {
	repo.lockClusterFile(path)
	defer repo.unlockClusterFile(path)

	_, status := os.Stat(path)
	fExists := os.IsNotExist(status)
	if fExists {
		return nil, nil
	}

	f, err := os.Open(path)
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

// overwriteClusterFile creates or updates the cluster by its path.
func (repo *cluster) overwriteClusterFile(path string, cluster *model.Cluster) error {
	repo.lockClusterFile(path)
	defer repo.unlockClusterFile(path)

	f, err := os.Create(path)
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

// assemblePath assembles the file cluster path using its name.
func (repo *cluster) assemblePath(name string) string {
	const F_EXTENSION = "gob"
	return fmt.Sprintf("%s/%s.%s", repo.basePath, name, F_EXTENSION)
}

// lockClusterFile locks the cluster file.
func (repo *cluster) lockClusterFile(path string) {
	smr, ok := repo.clusterSemaphore[path]
	if !ok {
		smr = &sync.Mutex{}
		repo.clusterSemaphore[path] = smr
	}
	smr.Lock()
}

// unlockClusterFile realizes the cluster file.
func (repo *cluster) unlockClusterFile(path string) {
	smr := repo.clusterSemaphore[path]
	smr.Unlock()
}
