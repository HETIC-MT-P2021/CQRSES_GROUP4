package pkg

import "reflect"

// TypeOf Get struct name
// ex : CreateArticleCommand{} => "CreateArticleCommand"
func TypeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}
