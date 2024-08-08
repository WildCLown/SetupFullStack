package services

import (
	"go-back/models"
)

func GetAllItems() []models.Item {
	return models.GetAllItems()
}

func CreateItem(item models.Item) {
	models.CreateItem(item)
}

// Outros métodos (GetItem, UpdateItem, DeleteItem) seguem o mesmo padrão
