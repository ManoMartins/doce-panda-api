package repository

import (
	"doce-panda/domain/order/entity"
	"doce-panda/domain/order/repository"
	productEntity "doce-panda/domain/product/entity"
	"doce-panda/infra/gorm/order/model"
	productModel "doce-panda/infra/gorm/product/model"
	"doce-panda/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type OrderRepositoryDb struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryDb {
	return &OrderRepositoryDb{Db: db}
}

func (o OrderRepositoryDb) FindById(ID string) (*entity.Order, error) {
	var orderModel model.Order

	o.Db.Preload("OrderItems").Preload("Address").Preload("User").Preload("OrderItems.Product").First(&orderModel, "id = ?", ID)

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
		User:         orderModel.User,
		Address:      orderModel.Address,
	})
}

func (o OrderRepositoryDb) FindAll() (*[]entity.Order, error) {
	var ordersModel []model.Order

	err := o.Db.Preload("OrderItems").Preload("OrderItems.Product").Order("created_at DESC").Find(&ordersModel).Error

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
				CreatedAt:    orderItemModel.CreatedAt,
				UpdatedAt:    orderItemModel.UpdatedAt,
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
			UserID:       orderModel.UserID,
			AddressID:    orderModel.AddressID,
			CreatedAt:    orderModel.CreatedAt,
			UpdatedAt:    orderModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}

	return &orders, nil
}

func (o OrderRepositoryDb) FindAllByUserId(UserID string) (*[]entity.Order, error) {
	var ordersModel []model.Order

	err := o.Db.Preload("OrderItems").Preload("OrderItems.Product").Order("created_at DESC").Find(&ordersModel, "user_id = ?", UserID).Error

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
			UserID:       orderModel.UserID,
			AddressID:    orderModel.AddressID,
			CreatedAt:    orderModel.CreatedAt,
			UpdatedAt:    orderModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		orders = append(orders, *order)
	}

	return &orders, nil
}

func (o OrderRepositoryDb) Create(order entity.Order) error {
	var paymentsModel []model.OrderPayment

	for _, pm := range order.Payments {
		paymentsModel = append(paymentsModel, model.OrderPayment{
			CreditCardID: pm.ID,
			OrderID:      order.ID,
			TotalInCents: pm.TotalInCents,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})
	}

	var couponsModel []model.OrderCoupon
	if len(order.Coupons) > 0 {
		for _, c := range order.Coupons {
			couponsModel = append(couponsModel, model.OrderCoupon{
				CouponID:  c.ID,
				OrderID:   order.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
		}
	}

	orderModel := model.Order{
		ID:           order.ID,
		TotalInCents: order.TotalInCents,
		Status:       order.Status,
		CouponID:     order.CouponID,
		AddressID:    order.AddressID,
		UserID:       order.UserID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := o.Db.Omit("Payments").Create(&orderModel).Error
	o.Db.Model(&orderModel).Association("Payments").Append(paymentsModel)
	o.Db.Model(&orderModel).Association("Coupons").Append(couponsModel)

	if err != nil {
		return err
	}

	randate := utils.Randate()

	var orderItemsModel []model.OrderItem
	for _, orderItem := range order.OrderItems {
		orderItemModel := model.OrderItem{
			ID:           orderItem.ID,
			ProductID:    orderItem.ProductID,
			OrderID:      orderModel.ID,
			Quantity:     orderItem.Quantity,
			TotalInCents: orderItem.TotalInCents,
			CreatedAt:    randate,
			UpdatedAt:    randate,
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

func (o OrderRepositoryDb) UpdateStatus(order entity.Order) error {
	var orderItemsModel []model.OrderItem

	for _, orderItem := range order.OrderItems {
		orderItemsModel = append(orderItemsModel, model.OrderItem{
			ID:           orderItem.ID,
			ProductID:    orderItem.ProductID,
			OrderID:      orderItem.OrderID,
			Quantity:     orderItem.Quantity,
			TotalInCents: orderItem.TotalInCents,
			Product: productModel.Product{
				ID:           orderItem.Product.ID,
				Name:         orderItem.Product.Name,
				PriceInCents: orderItem.Product.PriceInCents,
				Status:       orderItem.Product.Status,
				Description:  orderItem.Product.Description,
				Flavor:       orderItem.Product.Flavor,
				Quantity:     orderItem.Product.Quantity,
				ImageUrl:     orderItem.Product.ImageUrl,
				CreatedAt:    orderItem.Product.CreatedAt,
				UpdatedAt:    orderItem.Product.UpdatedAt,
			},
			CreatedAt: orderItem.CreatedAt,
			UpdatedAt: orderItem.UpdatedAt,
		})
	}

	var paymentsModel []model.OrderPayment

	for _, pm := range order.Payments {
		paymentsModel = append(paymentsModel, model.OrderPayment{
			CreditCardID: pm.ID,
			OrderID:      order.ID,
			TotalInCents: pm.TotalInCents,
			CreatedAt:    pm.CreatedAt,
			UpdatedAt:    pm.UpdatedAt,
		})
	}

	orderModel := model.Order{
		ID:     order.ID,
		Status: order.Status,
	}

	err := o.Db.Model(&orderModel).Update("status", orderModel.Status).Error

	if err != nil {
		return err
	}

	return nil
}

func (o OrderRepositoryDb) Report(input repository.InputReport) (*[]entity.OrderItem, error) {
	var orderItemsModel []model.OrderItem

	err := o.Db.Preload("Product").Preload("Product.Category").Where("created_at BETWEEN ? AND ?", input.StartDate, input.EndDate).Find(&orderItemsModel).Error

	if err != nil {
		return nil, err
	}

	var orderItems []entity.OrderItem
	for _, orderItemModel := range orderItemsModel {

		category, err := productEntity.NewCategory(productEntity.Category{
			ID:          orderItemModel.Product.Category.ID,
			Description: orderItemModel.Product.Category.Description,
			CreatedAt:   orderItemModel.Product.Category.CreatedAt,
			UpdatedAt:   orderItemModel.Product.Category.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		product, err := productEntity.NewProduct(productEntity.Product{
			ID:           orderItemModel.Product.ID,
			Name:         orderItemModel.Product.Name,
			PriceInCents: orderItemModel.Product.PriceInCents,
			Status:       orderItemModel.Product.Status,
			Description:  orderItemModel.Product.Description,
			Flavor:       orderItemModel.Product.Flavor,
			Quantity:     orderItemModel.Product.Quantity,
			Category:     *category,
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
			CreatedAt:    orderItemModel.CreatedAt,
			UpdatedAt:    orderItemModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, *orderItem)
	}

	return &orderItems, nil
}
