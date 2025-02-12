package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type task struct {
	gorm.Model
	User      string `json:"user"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func CreateTask(c *gin.Context) {
	var newTask task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	result := DB.Create(&newTask)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
}

func listTasks(c *gin.Context) {
	var tasks []task

	if err := DB.Find(&tasks).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	var target task

	if err := DB.First(&target, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, target)
}

func updateTaskByID(c *gin.Context) {
	id := c.Param("id")

	var newTask task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	var oldTask task
	if err := DB.First(&oldTask, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}

	oldTask.User = newTask.User
	oldTask.Task = newTask.Task
	oldTask.Completed = newTask.Completed

	if err := DB.Save(&oldTask).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusCreated, oldTask)
}

func deleteTaskByID(c *gin.Context) {
	id := c.Param("id")

	if err := DB.Delete(&task{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Tarefa Deletada!"})
}

func main() {

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao se conectar no banco de dados: ", err)
		return
	}
	DB = db
	err = DB.AutoMigrate(&task{})
	if err != nil {
		log.Fatal("Erro ao estruturar banco de dados: ", err)
		return
	}

	router := gin.Default()
	router.POST("todo/add_task", CreateTask)
	router.GET("todo/", listTasks)
	router.GET("todo/:id", getTaskByID)
	router.PUT("todo/:id", updateTaskByID)
	router.DELETE("todo/:id", deleteTaskByID)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Erro ao desponibilizar endpoint: ", err)
		return
	}
}
