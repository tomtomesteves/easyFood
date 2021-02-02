package models

import "easyfood/pkg/entity"

func NewCategory(c ...*entity.Category) []*Category {
	result := make([]*Category, 0)
	for _, category := range c {
		result = append(result, &Category{
			ID:   category.Id,
			Name: category.Name,
		})
	}

	return result
}
