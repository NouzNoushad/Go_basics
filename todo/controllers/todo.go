package controllers

import (
	"fmt"
	"gobasics/todo/config"
	"gobasics/todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// create todo
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := config.DB.Create(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// get todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo

	if result := config.DB.Order("updated_at desc").Find(&todos); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": fmt.Sprintf("%d items", len(todos)), "data": todos})
}

// get todo by id
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if result := config.DB.First(&todo, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, todo)
}

// update todo
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if result := config.DB.First(&todo, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := config.DB.Save(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// delete todo
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if result := config.DB.Find(&todo, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	if result := config.DB.Delete(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
