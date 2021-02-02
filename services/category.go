package services

import (
	"context"
	"easyfood/pkg/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CategoryService interface {
	Get(ctx context.Context, id *int) ([]*entity.Category, error)
	GetByDish(ctx context.Context, dishId int) (*entity.Category, error)
	GetByRestaurant(ctx context.Context, restaurantId int) ([]*entity.Category, error)
}

type categoryService struct {
	db *sqlx.DB
}

func NewCategoryService(db *sqlx.DB) CategoryService {
	return categoryService{db: db}
}

func (c categoryService) Get(ctx context.Context, id *int) ([]*entity.Category, error) {
	result := make([]*entity.Category,0)

	query := `SELECT * FROM categorias`

	if id != nil {
		query = fmt.Sprintf(`SELECT * FROM categorias WHERE id = %d`, *id)
	}
	err := c.db.Select(&result, query)

	if err != nil{
		return nil, err
	}
	return result, err
}

func (c categoryService) GetByDish(ctx context.Context, dishId int) (*entity.Category, error) {
	result := new(entity.Category)

	query := `SELECT * FROM categorias as c 
				INNER JOIN pratos as p
				on c.id = p.id_categoria
				WHERE p.id = ?`

	err := c.db.Get(result, query, dishId)
	if err != nil{
		return nil, err
	}
	return result, err
}

func (c categoryService) GetByRestaurant(ctx context.Context, restaurantId int) ([]*entity.Category, error) {
	result := make([]*entity.Category,0)

	query := `SELECT * FROM categorias as c 
				INNER JOIN restaurante-categoria as rc
				on c.id = rc.id_categoria
				WHERE rc.id_restaurante = ?`

	err := c.db.Select(&result, query, restaurantId)
	if err != nil{
		return nil, err
	}
	return result, err}



