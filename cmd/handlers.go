package main

import "siddharthkundu/pizza-shop-order/internal/models"

type Handler struct {
	orders *models.OrderModel
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		orders: &dbModel.Order,
	}
}
