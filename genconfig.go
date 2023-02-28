package loadconfig

import (
	"errors"
	"os"
	"reflect"
)

var (
	envTag      = "env"
	requiredTag = "required"
	defaultTag  = "default"
)

// GenConfig generates a config implementation from the config type defined
// in T.
func GenConfig[T any]() T {
	var cnf T
	v := reflect.Indirect(reflect.ValueOf(&cnf))
	t := reflect.TypeOf(cnf)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		varName := f.Tag.Get(envTag)

		val := os.Getenv(varName)
		if val == "" {
			if f.Tag.Get(requiredTag) == "true" {
				panic(errors.New("missing value for required parameter " + f.Name))
			} else {
				val = f.Tag.Get(defaultTag)
			}
		}
		v.FieldByName(f.Name).SetString(val)
	}
	return cnf
}
