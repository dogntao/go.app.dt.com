package controllers

type OrderController struct {
	BaseController
}

func (o *OrderController) Add() {
	o.ALoginCheck()
	if req.Method == "GET" {
		o.DisplayAdmin("views/order/add.html")
	} else {
	}
}
