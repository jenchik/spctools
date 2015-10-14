package spctools

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
)

func Caller(steps int) string {
	name := "?"
	if pc, _, _, ok := runtime.Caller(steps + 1); ok {
		name = filepath.Base(runtime.FuncForPC(pc).Name())
	}
	return name
}

func MakeBoundedIntFunc(minimum, maximum int, defaultValue *int) func(int) int {
	return func(x int) int {
		if x < 1 {
			x = *defaultValue
		}
		valid := x
		switch {
		case x < minimum:
			valid = minimum
		case x > maximum:
			valid = maximum
		}
		if valid != x {
			log.Printf("%s(): replaced %d with %d\n", Caller(1), x, valid)
		}
		return valid
	}
}

func GetField(obj interface{}, name string) (interface{}, error) {
    field := reflect.Indirect(reflect.ValueOf(obj)).FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("No such field: %s in obj", name)
	}
	return field.Interface(), nil
}

func GetName(obj interface{}) string {
    return reflect.Indirect(reflect.ValueOf(obj)).Type().Name()
}
