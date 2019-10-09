package main

import (
	"github.com/gin-gonic/gin"
)

type API struct {
	e          *gin.Engine
	Port       string
	repository RepositoryDB
	prefix     string
}

func newVis(repo RepositoryDB, prefix, port string) *API {
	return &API{
		e:          gin.Default(),
		Port:       port,
		prefix:     prefix,
		repository: repo,
	}
}

func (api *API) registerEndpoints() {

	r := api.e.Group(api.prefix)
	api.registerProjects(r)

}

func (api *API) Launch() error {
	api.registerEndpoints()
	return api.e.Run(api.Port)
}
