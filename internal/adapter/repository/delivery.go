package repository

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/adapter/repository/dbo"
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

type DeliveryRepository struct {
	db *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) *DeliveryRepository {
	return &DeliveryRepository{db}
}

func (d *DeliveryRepository) GetDeliveryById(id uint) (*entity.Delivery, error) {
	var delivery dbo.Delivery

	if err := d.db.Where("id = ?", id).First(&delivery).Error; err != nil {
		return nil, fmt.Errorf("error to find attendant with id %d - %v", id, err)
	}

	return delivery.ToEntity(), nil
}

func (d *DeliveryRepository) CreateDelivery(delivery *entity.Delivery) (*entity.Delivery, error) {
	deliveryEntity := dbo.ToDeliveryDBO(delivery)

	if err := d.db.Create(&deliveryEntity).Error; err != nil {
		return nil, err
	}

	return deliveryEntity.ToEntity(), nil
}

func (d *DeliveryRepository) UpdateDelivery(delivery *entity.Delivery) (*entity.Delivery, error) {
	deliveryEntity := dbo.ToDeliveryDBO(delivery)

	if err := d.db.Save(&deliveryEntity).Error; err != nil {
		return nil, err
	}

	return deliveryEntity.ToEntity(), nil
}
