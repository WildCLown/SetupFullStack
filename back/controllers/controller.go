package controllers

import (
	"go-back/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	item := models.GetItemByID(uint(id))
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item não encontrado"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"dale": "PumpIt"})
}

func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.CreateItem(newItem)
	c.JSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if models.UpdateItem(updatedItem) {
		c.JSON(http.StatusOK, updatedItem)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}

func DeleteItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if models.DeleteItem(uint(id)) {
		c.JSON(http.StatusOK, gin.H{"status": "Item deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}
