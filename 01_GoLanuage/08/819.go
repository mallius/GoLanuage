package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   int    "账号"
	Name string "姓名"
	Sex  bool   "性别"
}

func main() {
	u := user{100, "张三", false}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s(%s=%v)\n", f.Tag, f.Name, v.Field(i).Interface())
	}
}
