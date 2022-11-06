package repository

import (
	"doce-panda/domain/order/entity"
	productEntity "doce-panda/domain/product/entity"
	"doce-panda/infra/gorm/order/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type OrderRepositoryDb struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryDb {
	return &OrderRepositoryDb{Db: db}
}

func (o OrderRepositoryDb) FindById(ID string) (*entity.Order, error) {
	var orderModel model.Order

	o.Db.Preload("OrderItems").Preload("OrderItems.Product").First(&orderModel, "id = ?", ID)

	if orderModel.ID == "" {
		return nil, fmt.Errorf("Pedido não foi encontrado")
	}

	var orderItems []entity.OrderItem
	for _, orderItemModel := range orderModel.OrderItems {
		product, err := productEntity.NewProduct(productEntity.Product{
			ID:           orderItemModel.Product.ID,
			Name:         orderItemModel.Product.Name,
			PriceInCents: orderItemModel.Product.PriceInCents,
			Status:       orderItemModel.Product.Status,
			Description:  orderItemModel.Product.Description,
			Flavor:       orderItemModel.Product.Flavor,
			Quantity:     orderItemModel.Product.Quantity,
			ImageUrl:     orderItemModel.Product.ImageUrl,
			CreatedAt:    orderItemModel.Product.CreatedAt,
			UpdatedAt:    orderItemModel.Product.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		orderItem, err := entity.NewOrderItem(entity.OrderItem{
			ID:           orderItemModel.ID,
			ProductID:    orderItemModel.ProductID,
			OrderID:      orderItemModel.OrderID,
			Quantity:     orderItemModel.Quantity,
			TotalInCents: orderItemModel.TotalInCents,
			Product:      *product,
		})

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, *orderItem)
	}

	return entity.NewOrder(entity.Order{
		ID:           orderModel.ID,
		OrderItems:   orderItems,
		TotalInCents: orderModel.TotalInCents,
		Status:       orderModel.Status,
	})
}

func (o OrderRepositoryDb) FindAll() (*[]entity.Order, error) {
	var ordersModel []model.Order

	err := o.Db.Preload("OrderItems").Preload("OrderItems.Product").Find(&ordersModel).Error

	if err != nil {
		return nil, err
	}

	var orders []entity.Order

	for _, orderModel := range ordersModel {
		var orderItems []entity.OrderItem
		for _, orderItemModel := range orderModel.OrderItems {
			product, err := productEntity.NewProduct(productEntity.Product{
				ID:           orderItemModel.Product.ID,
				Name:         orderItemModel.Product.Name,
				PriceInCents: orderItemModel.Product.PriceInCents,
				Status:       orderItemModel.Product.Status,
				Description:  orderItemModel.Product.Description,
				Flavor:       orderItemModel.Product.Flavor,
				Quantity:     orderItemModel.Product.Quantity,
				ImageUrl:     orderItemModel.Product.ImageUrl,
				CreatedAt:    orderItemModel.Product.CreatedAt,
				UpdatedAt:    orderItemModel.Product.UpdatedAt,
			})

			if err != nil {
				return nil, err
			}

			orderItem, err := entity.NewOrderItem(entity.OrderItem{
				ID:           orderItemModel.ID,
				ProductID:    orderItemModel.ProductID,
				OrderID:      orderItemModel.OrderID,
				Quantity:     orderItemModel.Quantity,
				TotalInCents: orderItemModel.TotalInCents,
				Product:      *product,
			})

			if err != nil {
				return nil, err
			}

			orderItems = append(orderItems, *orderItem)
		}

		order, err := entity.NewOrder(entity.Order{
			ID:           orderModel.ID,
			OrderItems:   orderItems,
			TotalInCents: orderModel.TotalInCents,
			Status:       orderModel.Status,
		})

		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}

	return &orders, nil
}

func (o OrderRepositoryDb) Create(order entity.Order) error {
	orderModel := model.Order{
		ID:           order.ID,
		TotalInCents: order.TotalInCents,
		Status:       order.Status,
	}

	err := o.Db.Create(&orderModel).Error

	if err != nil {
		return err
	}

	var orderItemsModel []model.OrderItem
	for _, orderItem := range order.OrderItems {
		orderItemModel := model.OrderItem{
			ID:           orderItem.ID,
			ProductID:    orderItem.ProductID,
			OrderID:      orderModel.ID,
			Quantity:     orderItem.Quantity,
			TotalInCents: orderItem.TotalInCents,
		}

		orderItemsModel = append(orderItemsModel, orderItemModel)
	}

	for _, orderItemModel := range orderItemsModel {
		err = o.Db.Create(&orderItemModel).Error

		if err != nil {
			return err
		}
	}

	return nil
}
