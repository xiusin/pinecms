package tests

import (
	"github.com/go-xorm/xorm"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"reflect"
	"testing"
)

func TestYaegi(t *testing.T) {
	var err error
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	err = i.Use(interp.Exports{
		"pinecms/pinecms": {
			"Engine": reflect.ValueOf(&xorm.Engine{}),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = i.Eval(`
package tasks

import "pinecms"

func Run(orm pinecms.Engine) string {
	return "hello world"
}
`)
	if err != nil {
		t.Fatal(err)
	}
	v, err := i.Eval("tasks.Run")
	if err != nil {
		panic(err)
	}
	bar := v.Interface().(func(*xorm.Engine) string)
	r := bar(&xorm.Engine{})
	println(r)
}
