package structP

import (
	"fmt"
	"reflect"
	"strings"
)

func FmtStruct(st any, preName ...string) {
	vo := reflect.ValueOf(st)
	t := vo.Type()
	if vo.Kind() == reflect.Ptr {
		vo = vo.Elem()
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("not struct")
		return
	}
	for i := 0; i < t.NumField(); i++ {
		if vo.Field(i).Kind() == reflect.Struct {
			newPreName := append(preName, t.Field(i).Name)
			FmtStruct(vo.Field(i).Interface(), newPreName...)
			continue
		}
		names := append(preName, t.Field(i).Name)
		name := strings.Join(names, ".")
		fmt.Printf("%s: %v\n", name, vo.Field(i).Interface())
	}
}
