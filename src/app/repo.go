package app

import "github.com/viniciuscrisol/dynamic-db-v2/app/model"

type ClusterRepo interface {
	// AssembleURL assembles the cluster URL using the cluster name.
	AssembleURL(name string) string

	// Create stores a cluster in the repository.
	Create(cluster *model.Cluster) error

	// Update updates a stored cluster in the repository by its ID and URL.
	Update(cluster *model.Cluster) error

	// Find searches for a cluster by its URL. If a URL matches, it will be returned.
	// Otherwise, a nil pointer will be returned.
	Find(url string) (*model.Cluster, error)
}
