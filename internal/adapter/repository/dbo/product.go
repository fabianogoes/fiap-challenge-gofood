package dbo

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

// Product is a Database Object for product
type Product struct {
	gorm.Model
	Name       string `gorm:"unique"`
	Price      float64
	CategoryID uint
	Category   Category `gorm:"ForeignKey:CategoryID"`
}

func (p *Product) ToEntity() *entity.Product {
	return &entity.Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
		Category: &entity.Category{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (p *Product) ToDBO(product *entity.Product) *Product {
	return &Product{
		Model: gorm.Model{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		},
		Name:       product.Name,
		Price:      product.Price,
		CategoryID: product.Category.ID,
	}
}

// Category is a Database Object for category
type Category struct {
	gorm.Model
	Name string
}

func (c *Category) ToModel() *entity.Category {
	return &entity.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (c *Category) ToDBO(category *entity.Category) *Category {
	return &Category{
		Model: gorm.Model{
			ID:        category.ID,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
		Name: category.Name,
	}
}
