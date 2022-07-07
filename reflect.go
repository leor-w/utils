package utils

import "reflect"

func IsNilPointer(i interface{}) bool {
	iv := reflect.ValueOf(i)
	return iv.IsNil() || iv.Kind() != reflect.Ptr
}

func IsFunc(i interface{}) bool {
	iv := reflect.ValueOf(i)
	return !iv.IsNil() && iv.Kind() == reflect.Func
}

func RemoveTypePtr(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func RemoveValuePtr(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
