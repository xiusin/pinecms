package backend

import (
	"fmt"
	"github.com/xiusin/pine"
	"reflect"
)

type BaseController struct {
	I
	pine.Controller
}

type I interface {
	Construct()
}

func (c *BaseController) List() {
	fmt.Println(reflect.TypeOf(c).Elem())
}

func (c *BaseController) Add() {

}

func (c *BaseController) Edit() {

}

func (c *BaseController) Order() {

}

func (c *BaseController) Delete() {

}

func (c *BaseController) Info() {

}
