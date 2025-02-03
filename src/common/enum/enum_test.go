package enum

import "testing"

type Ttt struct {
	CustomValue Type[string] `enum:"name_custom_value"`
	KeyAsValue  Type[string]
}

func TestEnum(t *testing.T) {
	New[Ttt](nil)
}
