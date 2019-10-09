package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) registerProjects(r *gin.RouterGroup) {
	r.POST("/createProject", func(c *gin.Context) {
		name := c.Query("name")
		totalactions := c.Query("_total_actions")
		status := c.Query("status")
		project := &Project{
			Name:         name,
			TotalActions: totalactions,
			Status:       status,
		}
		err := api.repository.CreateNewProject(project)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": project,
		})
	})

	r.GET("/getProject/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "all" || id == "" {
			project, err := api.repository.Projects()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": project,
			})
		} else {
			project, err := api.repository.GetProjectByID(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"data": project,
			})
		}
	})

	r.PUT("/updateProject/:id", func(c *gin.Context) {
		project := new(Project)
		if err := c.BindJSON(project); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err := api.repository.UpdateProjectByID(project)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": project,
		})
	})

}
