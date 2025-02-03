// Package enum 枚举值类型
package enum

import (
	"fmt"
	"reflect"
	"strings"
)

// TypeContract 枚举值类型接口
type TypeContract interface {
	SetEnumValue(value any)
	String() string
}

// Type 枚举值类型
type Type[T any] struct {
	// ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	// 	~float32 | ~float64 | ~string | ~bool
	value *T
}

// SetEnumValue 设置枚举值类型
func (t Type[T]) SetEnumValue(value any) {
	t.value = value.(*T)
	fmt.Printf("%p\n", t.value)

}

func (t Type[T]) String() string {
	return fmt.Sprintf("%v", t.value)
}

// New 枚举值类型
func New[T any](t *T) *T {
	if t == nil {
		t = new(T)
	}
	value := reflect.ValueOf(t).Elem()
	typ := reflect.TypeOf(t).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fv := value.Field(i)

		// 非枚举定类型忽略
		if !strings.Contains(field.Type.String(), "enum.Type") {
			continue
		}

		// 判断具体反射类型
		valueType, ok := field.Type.FieldByName("value")
		if !ok {
			continue
		}

		// 转换为枚举类型定义
		//  fv.Call()
		// fmt.Println(any(&fvi).(TypeContract))

		// fvt, ok := (any(fv.Interface())).(TypeContract)
		// if !ok {
		// 	continue
		// }

		method := fv.MethodByName("SetEnumValue") // 假设你要调用的方法名是 SetEnumValue
		if !method.IsValid() {
			continue
		}

		fieldValue := field.Tag.Get("enum")
		switch valueType.Type.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// fvt.SetEnumValue(&i)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			// fvt.SetEnumValue(&i)
		case reflect.String:
			if fieldValue == "" {
				fieldValue = field.Name
			}
			method.Call([]reflect.Value{reflect.ValueOf(&fieldValue)})

			fmt.Println(t)
		default:

			panic("enum value type not support: " + valueType.Type.String())
		}

	}
	return t

}
