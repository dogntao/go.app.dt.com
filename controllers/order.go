package controllers

type OrderController struct {
	BaseController
}

func (o *OrderController) Add() {
	if req.Method == "GET" {
		o.Display("views/order/add.html")
	} else {
	}
}
