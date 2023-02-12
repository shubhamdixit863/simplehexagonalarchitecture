package repository

import "task4/cmd/web/models"

type Repository interface {
	GetItems() ([]models.Item, error)
	GetItemById(id string) (*models.Item, error)
	AddItems(item models.Item) error
	UpdateItem(id string, item models.Item) error
	DeleteItem(id string) error
}
