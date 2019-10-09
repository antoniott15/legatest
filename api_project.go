package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) registerProjects(r *gin.RouterGroup) {

	r.POST("/createProject", func(c *gin.Context) {
		name := c.Query("name")
		totalActions := c.Query("totalactions")
		status := c.Query("status")
		location := c.Query("location")

		project := &Project{
			Name:         name,
			TotalActions: totalActions,
			Status:       status,
			Location:     location,
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
		name := c.Query("name")
		totalActions := c.Query("totalactions")
		status := c.Query("status")
		id := c.Param("id")
		location := c.Query("location")

		lastUpdate, err := api.repository.GetProjectByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if name == "" {
			name = lastUpdate.Name
		}
		if totalActions == "" {
			totalActions = lastUpdate.TotalActions
		}
		if status == "" {
			status = lastUpdate.Status
		}
		if location == "" {
			location = lastUpdate.Location
		}
		creationdate := lastUpdate.CreationDate
		project := &Project{
			ID:           id,
			Name:         name,
			TotalActions: totalActions,
			Status:       status,
			Location:     location,
			CreationDate: creationdate,
		}
		err = api.repository.UpdateProjectByID(project)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": project,
		})
	})

}
