package controllers

type CustomerController struct {
	BaseController
}

func (c *CustomerController) Add() {
	if req.Method == "GET" {
		c.Display("views/customer/add.html")
	} else {
	}
}

func (c *CustomerController) Manage() {
	if req.Method == "GET" {
		c.Display("views/customer/manage.html")
	} else {
	}
}
