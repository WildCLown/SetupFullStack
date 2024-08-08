package models

type Item struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item
var nextID uint = 1

func GetAllItems() []Item {
	return items
}

func CreateItem(item Item) Item {
	item.ID = nextID
	nextID++
	items = append(items, item)
	return item
}

func GetItemByID(id uint) *Item {
	for _, item := range items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func UpdateItem(updatedItem Item) bool {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i] = updatedItem
			return true
		}
	}
	return false
}

func DeleteItem(id uint) bool {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return true
		}
	}
	return false
}
