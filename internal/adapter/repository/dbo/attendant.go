package dbo

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

// Attendant is a Database Object for attendant
type Attendant struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// ToEntity converts Attendant DBO to entity.Attendant
func (a *Attendant) ToEntity() *entity.Attendant {
	return &entity.Attendant{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Name:      a.Name,
	}
}

// ToDBO converts entity.Attendant to Attendant DBO
func ToAttendantDBO(a *entity.Attendant) *Attendant {
	return &Attendant{
		Model: gorm.Model{
			ID:        a.ID,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
		},
		Name: a.Name,
	}
}
