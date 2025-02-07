package enum

import (
	"fmt"
	"testing"
)

type Ttt struct {
	CustomValue Type[string] `enum:"name_custom_value"`
	KeyAsValue  Type[string]
}

var vv = New[Ttt](nil)

func TestEnum(_ *testing.T) {
	EnumAsParam(vv.CustomValue)
}

func EnumAsParam(dist TypeContract[string]) {
	fmt.Println(dist.Value())
}
