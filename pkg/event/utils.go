package event

import (
	"reflect"
)

func typeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}
