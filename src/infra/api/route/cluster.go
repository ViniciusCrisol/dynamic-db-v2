package route

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciuscrisol/dynamic-db-v2/app/usecase"
	"github.com/viniciuscrisol/dynamic-db-v2/infra/api"
)

type cluster struct {
	createCluster *usecase.CreateCluster
}

// NewCluster returns a cluster router.
func NewCluster(createCluster *usecase.CreateCluster) *cluster {
	return &cluster{createCluster}
}

// Create handles an HTTP request to create a cluster and return it to the client.
func (rtr *cluster) Create(ctx *gin.Context) {
	b := &api.CreateClusterDTO{}
	err := api.BindRequestBody(b, ctx)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}

	c, err := rtr.createCluster.Exec(b.Name)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}
	api.SendJSON(200, c, ctx)
}
