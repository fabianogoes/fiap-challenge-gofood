package repository

import (
	"fmt"
	"time"

	"github.com/fiap/challenge-gofood/internal/adapter/repository/dbo"
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db             *gorm.DB
	itemRepository *OrderItemRepository
}

func NewOrderRepository(db *gorm.DB, itemRepo *OrderItemRepository) *OrderRepository {
	return &OrderRepository{db, itemRepo}
}

func (or *OrderRepository) CreateOrder(entity *domain.Order) (*domain.Order, error) {
	order := &dbo.Order{
		CustomerID:  entity.Customer.ID,
		AttendantID: entity.Attendant.ID,
		Date:        time.Now(),
		Status:      entity.Status.ToString(),
		Payment:     dbo.ToPaymentDBO(entity.Payment),
		Delivery:    dbo.ToDeliveryDBO(entity.Delivery),
		Amount:      entity.Amount(),
	}

	if err := or.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order.ToEntity(), nil
}

func (or *OrderRepository) GetOrderById(id uint) (*domain.Order, error) {
	order := &dbo.Order{}

	if err := or.db.Preload("Customer").
		Preload("Attendant").
		Preload("Payment").
		Preload("Delivery").
		Preload("Items").
		First(order, id).Error; err != nil {
		return nil, fmt.Errorf("error to find order with id %d - %v", id, err)
	}

	for _, item := range order.Items {
		product := &dbo.Product{}
		if err := or.db.First(product, item.ProductID).Error; err != nil {
			return nil, fmt.Errorf("error to find product with id %d - %v", item.ProductID, err)
		}
		item.Product = product
	}

	return order.ToEntity(), nil
}

func (or *OrderRepository) UpdateOrder(order *domain.Order) (*domain.Order, error) {
	orderToUpdate := &dbo.Order{}

	if err := or.db.Preload("Customer").Preload("Attendant").Preload("Payment").Preload("Items").
		First(orderToUpdate, order.ID).Error; err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		if item.ID == 0 {
			if err := or.itemRepository.CreateOrderItem(dbo.ToOrderItemDBO(item)); err != nil {
				return nil, err
			}
		}
	}

	orderToUpdate.Amount = order.Amount()
	orderToUpdate.Status = order.Status.ToString()
	orderToUpdate.Payment = dbo.ToPaymentDBO(order.Payment)

	if err := or.db.Save(orderToUpdate).Error; err != nil {
		return nil, err
	}

	return or.GetOrderById(order.ID)
}

func (or *OrderRepository) RemoveItemFromOrder(idItem uint) error {
	return or.itemRepository.Delete(idItem)
}

func (or *OrderRepository) GetOrderItemById(id uint) (*domain.OrderItem, error) {
	return or.itemRepository.GetOrderItemById(id)
}

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db}
}

func (oir *OrderItemRepository) CreateOrderItem(orderItem *dbo.OrderItem) error {
	if err := oir.db.Create(orderItem).Error; err != nil {
		return err
	}

	return nil
}

func (oir *OrderItemRepository) Delete(idItem uint) error {
	if err := oir.db.Delete(&dbo.OrderItem{}, idItem).Error; err != nil {
		return err
	}

	return nil
}

func (oir *OrderItemRepository) GetOrderItemById(id uint) (*domain.OrderItem, error) {
	orderItem := &dbo.OrderItem{}

	if err := oir.db.Preload("Product").First(orderItem, id).Error; err != nil {
		return nil, fmt.Errorf("error to find order item with id %d - %v", id, err)
	}

	return orderItem.ToEntity(), nil
}
