package main

import "reflect"

/*
	golang challenge: write a function walk(x interface{},fn func(string))
	which takes a struct x and calls fn for all strings fields found inside.
	difficulty level: recursively

*/

func walk(x interface {}, fn func(string)){
	val := getValue(x)

	walkValue := func(value reflect.Value){
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++{
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++{
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys(){
			walk(val.MapIndex(key).Interface(),fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value{
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr{
		val = val.Elem()
	}

	return val
}



