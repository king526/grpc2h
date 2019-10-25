package grpc2h

import (
	"reflect"
)

func setFieldValueByType(rs reflect.Value, fieldType reflect.Type, val interface{}) {
	// get struct
	for rs.Type().Kind() != reflect.Struct {
		rs = rs.Elem()
	}
	for i := 0; i < rs.NumField(); i++ {
		if rf := rs.Field(i); rf.Type() == fieldType {
			rf.Set(reflect.ValueOf(val))
		}
	}
}

func setFieldValueByName(rs reflect.Value, fieldName string, val interface{}) {
	// get struct
	for rs.Type().Kind() != reflect.Struct {
		rs = rs.Elem()
	}
	rs.FieldByName(fieldName).Set(reflect.ValueOf(val))
}
