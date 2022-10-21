package route

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciuscrisol/dynamic-db-v2/app/usecase"
	"github.com/viniciuscrisol/dynamic-db-v2/infra/api"
)

type schema struct {
	create     *usecase.CreateSchema
	deleteByID *usecase.DeleteSchemaByID
}

// NewSchema returns a schema router.
func NewSchema(create *usecase.CreateSchema, deleteByID *usecase.DeleteSchemaByID) *schema {
	return &schema{
		create:     create,
		deleteByID: deleteByID,
	}
}

// Create handles an HTTP request to create a schema and return it to the client.
func (rtr *schema) Create(ctx *gin.Context) {
	b := &api.CreateSchemaDTO{}
	err := api.BindRequestBody(b, ctx)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}

	c, err := rtr.create.Exec(b.Name, b.Content)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}
	api.SendJSON(200, c, ctx)
}

// DeleteByID handles an HTTP request to delete a schema by ID.
func (rtr *schema) DeleteByID(ctx *gin.Context) {
	name := ctx.Param("name")
	id := ctx.Param("id")

	err := rtr.deleteByID.Exec(name, id)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}
	api.SendJSON(201, nil, ctx)
}
