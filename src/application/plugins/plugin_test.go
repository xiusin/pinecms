package plugins

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"io/ioutil"
	"testing"
)

func TestYaegi(t *testing.T) {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	a,_ := ioutil.ReadFile("/Users/xiusin/projects/src/github.com/xiusin/pinecms/src/application/plugins/a.yaegi")


	t.Log(i.Eval(string(a)))

}
