package pkg

import (
	"reflect"
)

func TypeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}
