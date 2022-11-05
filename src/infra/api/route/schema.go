package route

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciuscrisol/dynamic-db-v2/app/usecase"
	"github.com/viniciuscrisol/dynamic-db-v2/infra/api"
)

type schema struct {
	create          *usecase.CreateSchema
	deleteByID      *usecase.DeleteSchemaByID
	deleteByContent *usecase.DeleteSchemaByContent
	filter          *usecase.FilterSchemas
}

// NewSchema returns a schema router.
func NewSchema(
	create *usecase.CreateSchema,
	deleteByID *usecase.DeleteSchemaByID,
	deleteByContent *usecase.DeleteSchemaByContent,
	filter *usecase.FilterSchemas,
) *schema {
	return &schema{
		create:          create,
		deleteByID:      deleteByID,
		deleteByContent: deleteByContent,
		filter:          filter,
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
	name := ctx.Param("name")

	c, err := rtr.create.Exec(name, b.Content)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}
	api.SendJSON(201, c, ctx)
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
	api.SendJSON(204, nil, ctx)
}

// DeleteByContent handles an HTTP request to delete schemas by it content.
func (rtr *schema) DeleteByContent(ctx *gin.Context) {
	name := ctx.Param("name")
	content := map[string]string{}
	for k, v := range ctx.Request.URL.Query() {
		content[k] = v[0]
	}

	err := rtr.deleteByContent.Exec(name, content)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}
	api.SendJSON(204, nil, ctx)
}

// Filter handles an HTTP request to filter schemas by it content.
func (rtr *schema) Filter(ctx *gin.Context) {
	name := ctx.Param("name")
	content := map[string]string{}
	for k, v := range ctx.Request.URL.Query() {
		content[k] = v[0]
	}

	schemas, err := rtr.filter.Exec(name, content)
	if err != nil {
		api.HandleErr(err, ctx)
		return
	}
	api.SendJSON(200, schemas, ctx)
}
