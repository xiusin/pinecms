package enum

import (
	"fmt"
	"testing"
)

type Ttt struct {
	CustomValue Type[string] `enum:"name_custom_value"`
	KeyAsValue  Type[string]
}

func TestEnum(t *testing.T) {
	vv := New[Ttt](nil)

	fmt.Println(vv.CustomValue)
}
