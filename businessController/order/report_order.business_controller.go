package order

import (
	"doce-panda/businessController/order/dtos"
	"doce-panda/domain/order/repository"
)

type ReportOrderBusinessController struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewReportOrderBusinessController(orderRepository repository.OrderRepositoryInterface) *ReportOrderBusinessController {
	return &ReportOrderBusinessController{orderRepository: orderRepository}
}

func (r ReportOrderBusinessController) Execute(input dtos.InputReportOrderDto) (*dtos.OutputReportOrderDto, error) {
	orders, err := r.orderRepository.Report(repository.InputReport{
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	})

	if err != nil {
		return nil, err
	}

	mapOrder := make(map[string][]dtos.Order)
	for _, order := range *orders {
		mapOrder[order.Product.Category.ID] = append(mapOrder[order.Product.Category.ID], dtos.Order{
			CategoryName: order.Product.Category.Description,
			PriceInCents: order.TotalInCents,
			CreatedAt:    order.CreatedAt,
		})
	}

	var series []dtos.Series
	for key, value := range mapOrder {
		series = append(series, dtos.Series{
			CategoryID:   key,
			CategoryName: value[0].CategoryName,
			Orders:       value,
		})
	}

	return &dtos.OutputReportOrderDto{
		Series:    series,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		//StartDate: time.Date(2023, 2, 1, 0, 0, 0, 0, time.Local).Format(time.RFC3339),
		//EndDate:   time.Date(2023, 2, 28, 0, 0, 0, 0, time.Local).Format(time.RFC3339),
	}, nil
}
