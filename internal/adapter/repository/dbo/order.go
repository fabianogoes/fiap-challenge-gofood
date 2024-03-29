package dbo

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

// Order is a Database Object for order
type Order struct {
	gorm.Model
	CustomerID  uint
	Customer    Customer `gorm:"ForeignKey:CustomerID"`
	AttendantID uint
	Attendant   Attendant `gorm:"ForeignKey:AttendantID"`
	Date        time.Time
	Status      string
	PaymentID   uint
	Payment     Payment `gorm:"ForeignKey:PaymentID"`
	Amount      float64
	DeliveryID  uint
	Delivery    Delivery `gorm:"ForeignKey:DeliveryID"`
	Items       []*OrderItem
}

// ToEntity converts Order DBO to domain.Order
func (o *Order) ToEntity() *domain.Order {
	var items []*domain.OrderItem
	var itemsTotal int

	for _, item := range o.Items {
		items = append(items, item.ToEntity())
		itemsTotal += int(item.Quantity)
	}

	return &domain.Order{
		ID: o.ID,
		Customer: &domain.Customer{
			ID:   o.Customer.ID,
			Name: o.Customer.Name,
			CPF:  o.Customer.CPF,
		},
		Attendant: &domain.Attendant{
			ID:   o.Attendant.ID,
			Name: o.Attendant.Name,
		},
		Date:      o.Date,
		Status:    o.toOrderStatus(),
		Payment:   o.Payment.ToEntity(),
		Delivery:  o.Delivery.ToEntity(),
		Items:     items,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

// ToDBO converts domain.Order to Order DBO
func (o *Order) toOrderStatus() domain.OrderStatus {
	switch o.Status {
	case "STARTED":
		return domain.OrderStatusStarted
	case "ADDING_ITEMS":
		return domain.OrderStatusAddingItems
	case "CONFIRMED":
		return domain.OrderStatusConfirmed
	case "PAID":
		return domain.OrderStatusPaid
	case "PAYMENT_REVERSED":
		return domain.OrderStatusPaymentReversed
	case "IN_PREPARATION":
		return domain.OrderStatusInPreparation
	case "READY_FOR_DELIVERY":
		return domain.OrderStatusReadyForDelivery
	case "SENT_FOR_DELIVERY":
		return domain.OrderStatusSentForDelivery
	case "DELIVERED":
		return domain.OrderStatusDelivered
	case "CANCELED":
		return domain.OrderStatusCanceled
	default:
		return domain.OrderStatusStarted
	}
}

// ToDBO converts domain.Order to Order DBO
type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     Order
	ProductID uint
	Product   *Product
	Quantity  int
	UnitPrice float64
}

// ToEntity converts OrderItem DBO to domain.OrderItem
func (i *OrderItem) ToEntity() *domain.OrderItem {
	return &domain.OrderItem{
		ID:        i.ID,
		Product:   i.Product.ToEntity(),
		Quantity:  int(i.Quantity),
		UnitPrice: i.UnitPrice,
	}
}

// ToDBO converts domain.OrderItem to OrderItem DBO
func ToOrderItemDBO(i *domain.OrderItem) *OrderItem {
	return &OrderItem{
		Model: gorm.Model{
			ID:        i.ID,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
		},
		OrderID:   i.Order.ID,
		ProductID: i.Product.ID,
		Quantity:  i.Quantity,
		UnitPrice: i.UnitPrice,
	}
}
