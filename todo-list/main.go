package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var tasks = []task{
	{ID: "1", Task: "Ola Mundo! Tarefa 1"},
	{ID: "2", Task: "Ola Mundo! Tarefa 2"},
	{ID: "3", Task: "Ola Mundo! Tarefa 3"},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postTask(c *gin.Context) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func main() {
	router := gin.Default()
	router.GET("/todo/list", getTasks)
	router.POST("/todo/add_task", postTask)

	router.Run(":8080")
}
