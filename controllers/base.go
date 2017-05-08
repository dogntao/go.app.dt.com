package controllers

import (
	"fmt"
)

//
type BaseController struct {
}

func (ba *BaseController) Display() {
	fmt.Println("display")
}
