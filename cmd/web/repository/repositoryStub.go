package repository

import (
	"errors"
	"task4/cmd/web/models"
)

type Stub struct {
	Items []models.Item
}

func (r *Stub) GetItemById(id string) (*models.Item, error) {
	for _, item := range r.Items {
		if item.Id == id {

			return &item, nil
		}
	}
	return nil, errors.New("item Not Found")

}

func (r *Stub) AddItems(item models.Item) error {
	r.Items = append(r.Items, item)
	return nil
}

func (r *Stub) UpdateItem(id string, item models.Item) error {
	for index, item := range r.Items {
		if item.Id == id {
			r.Items[index] = item
			return nil
		}
	}
	return errors.New("item Not Found")
}

func (r *Stub) DeleteItem(id string) error {

	for index, item := range r.Items {
		if item.Id == id {
			r.Items = append(r.Items[:index], r.Items[index+1:]...)
			return nil

		}

	}
	return errors.New("item Not Found")
}

func (r *Stub) GetItems() ([]models.Item, error) {
	return r.Items, nil
}
