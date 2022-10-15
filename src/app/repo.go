package app

import "github.com/viniciuscrisol/dynamic-db-v2/app/model"

type ClusterRepo interface {
	// Create stores a cluster in the repository.
	Create(cluster *model.Cluster) error

	// Update updates a stored cluster in the repository by its name.
	Update(cluster *model.Cluster) error

	// Find searches for a cluster by its name. If a name matches, it will be returned.
	// Otherwise, a nil pointer will be returned.
	Find(name string) (*model.Cluster, error)
}
