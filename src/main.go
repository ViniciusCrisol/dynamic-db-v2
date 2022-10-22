package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/viniciuscrisol/dynamic-db-v2/app"
	"github.com/viniciuscrisol/dynamic-db-v2/app/usecase"
	"github.com/viniciuscrisol/dynamic-db-v2/infra/api"
	"github.com/viniciuscrisol/dynamic-db-v2/infra/api/route"
	"github.com/viniciuscrisol/dynamic-db-v2/infra/repo"
)

func main() {
	app.LoadEnv(".env")

	router := getRouter()
	clusterGroup := router.Group(app.HTTP_ROUTES_PREFIX + "/cluster")
	schemaGroup := router.Group(app.HTTP_ROUTES_PREFIX + "/cluster/:name/schema")
	// Dependencies
	repo := repo.NewCluster(app.STORAGE_PATH)
	// Cluster
	createCluster := usecase.NewCreateCluster(repo)
	clusterRouter := route.NewCluster(createCluster)
	clusterGroup.POST("/", clusterRouter.Create)
	// Schema
	createSchema := usecase.NewCreateSchema(repo)
	deleteSchemaByID := usecase.NewDeleteSchemaByID(repo)
	schemaRouter := route.NewSchema(
		createSchema,
		deleteSchemaByID,
	)
	schemaGroup.POST("/", schemaRouter.Create)
	schemaGroup.DELETE("/:id", schemaRouter.DeleteByID)

	initServer(router)
}

func getRouter() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	cps := gzip.Gzip(gzip.BestCompression)

	router := gin.Default()
	router.Use(cps)
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:  []string{"*"},
				AllowMethods:  []string{"*"},
				AllowHeaders:  []string{"*"},
				AllowWildcard: true,
			},
		),
	)
	router.NoRoute(api.SendRouteNotFound)
	return router
}

func initServer(router *gin.Engine) {
	err := router.Run(":" + app.WEB_SERVER_PORT)
	if err != nil {
		panic(err)
	}
}
