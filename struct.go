package aknstruct

import (
	"reflect"
	"strings"
)

func Map(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	val := reflect.ValueOf(s)
	for i := 0; i < val.Type().NumField(); i++ {
		value := reflect.Indirect(val).Field(i)
		// m[val.Type().Field(i).Name] = value.Interface()
		cast(m, val.Type().Field(i), value.Interface())
	}
	return m
}

func cast(m map[string]interface{}, field reflect.StructField, value interface{}) {
	name := field.Name
	tag := mapTag(field.Tag)
	if len(tag) > 0 {
		name = tag
	}
	if strings.Compare(tag, "-") != 0 {
		m[name] = value
	}
}

func mapTag(tag reflect.StructTag) string {
	tags := strings.Split(tag.Get("map"), ",")
	return tags[0]
}
