package test

import (
	"testing"

	"github.com/flosch/pongo2"
)

func Test_Pongo(t *testing.T)  {
	pongo2.RegisterTag("list", tagListTagParser)
	tpl ,err := pongo2.FromString("'{% for item in list %} hello world {% endfor %}'")
	if err != nil {
		t.Fatal(err)
	}
	s, err := tpl.Execute(pongo2.Context{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(s))
}