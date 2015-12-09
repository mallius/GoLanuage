package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Num   int
	Name  string
	Sex   bool
	Score float32
}

func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("Cannot set!")
	} else {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Int:
			v.Field(i).SetInt(10002)
		case reflect.String:
			v.Field(i).SetString("张恒")
		case reflect.Bool:
			v.Field(i).SetBool(true)
		case reflect.Float32:
			v.Field(i).SetFloat(95.5)
		}
	}
}
func main() {
	var stu student = student{1, "stu01", true, 100}
	fmt.Println(stu)
	SetValue(&stu)
	fmt.Println(stu)

}
