// Package enum 枚举值类型
package enum

import (
	"reflect"
	"strings"
	"unsafe"
)

// TypeContract 枚举值类型接口
type TypeContract[T any] interface {
	Value() T
}

// Type 枚举值类型
type Type[T any] struct {
	value *T
}

func (t Type[T]) Value() T {
	return *t.value
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

		rf := reflect.Indirect(fv).FieldByName("value")
		ptr := unsafe.Pointer(rf.UnsafeAddr())

		fieldValue := field.Tag.Get("enum")
		switch valueType.Type.Elem().Kind() {
		case reflect.Int:
			*(**int)(ptr) = toPtr(i + 1)
		case reflect.Int8:
			*(**int8)(ptr) = toPtr(int8(i + 1))
		case reflect.Int16:
			*(**int16)(ptr) = toPtr(int16(i + 1))
		case reflect.Int32:
			*(**int32)(ptr) = toPtr(int32(i + 1))
		case reflect.Int64:
			*(**int64)(ptr) = toPtr(int64(i + 1))
		case reflect.Uint:
			*(**uint)(ptr) = toPtr(uint(i + 1))
		case reflect.Uint8:
			*(**uint8)(ptr) = toPtr(uint8(i + 1))
		case reflect.Uint16:
			*(**uint16)(ptr) = toPtr(uint16(i + 1))
		case reflect.Uint32:
			*(**uint32)(ptr) = toPtr(uint32(i + 1))
		case reflect.Uint64:
			*(**uint64)(ptr) = toPtr(uint64(i + 1))
		case reflect.String:
			if fieldValue == "" {
				fieldValue = field.Name
			}
			*(**string)(ptr) = toPtr(fieldValue)
		default:
			panic("enum value type not support: " + valueType.Type.String())
		}
	}
	return t

}

func toPtr[T any](v T) *T {
	return &v
}
