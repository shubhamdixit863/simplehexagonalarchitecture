package dto

import "task4/cmd/web/models"

type ItemRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}


func (ir *ItemRequest) ToModel() models.Item{
	return models.Item{
		Id: ir.Id,
		Name: ir.Name,
		Description: ir.Description,
		Price: ir.Price,


	}

}