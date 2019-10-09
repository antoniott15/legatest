package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func (api *API) testingGraphql(r *gin.RouterGroup) {

	var projectType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Project",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"location": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"totalactions": &graphql.Field{
					Type: graphql.String,
				},
				"creationdate": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var queryProject = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"product": &graphql.Field{
					Type:        projectType,
					Description: "Get project by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(string)
						if ok {
							c, err := api.repository.GetProjectByID(id)
							if err != nil {
								log.Println(err)
							}
							return c, nil
						}
						return nil, nil
					},
				},
				"list": &graphql.Field{
					Type:        graphql.NewList(projectType),
					Description: "Get projects list",
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						c, err := api.repository.Projects()
						if err != nil {
							log.Println(err)
						}
						return c, nil
					},
				},
			},
		})

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        projectType,
				Description: "Create a new project",
				Args: graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"status": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"totalactions": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name := params.Args["name"].(string)
					totalactions := params.Args["totalactions"].(string)
					status := params.Args["status"].(string)
					location := params.Args["location"].(string)
					project := &Project{
						Name:         name,
						TotalActions: totalactions,
						Status:       status,
						Location:     location,
					}
					err := api.repository.CreateNewProject(project)
					if err != nil {
						log.Println(err)
					}
					return project, nil
				},
			},
		},
	})

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryProject,
			Mutation: mutationType,
		},
	)

	r.POST("/graphql", func(c *gin.Context) {
		query := c.Query("query")

		result := executeQuery(query, schema)

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})
}
