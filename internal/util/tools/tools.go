package tools

import (
	"errors"
	"fmt"
	"reflect"
	"search_engine/internal/service/objs"
	"strings"
	"time"
	"unsafe"
)

func TimeCost() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}

func Str2Bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2Str(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func Camel2SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func Snake2CamelString(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	return string(data[:])
}

func ConvStruct2Map(s interface{}) (map[string]objs.FieldInfo, error) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("kind is not struct")
	}

	fieldMap := make(map[string]objs.FieldInfo, 0)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct && v.Type().Field(i).Type.Name() == v.Type().Field(i).Name {
			if subFieldMap, err := ConvStruct2Map(v.Field(i).Interface()); err == nil {
				for fieldName, fieldInfo := range subFieldMap {
					fieldMap[fieldName] = fieldInfo
				}
			}
		} else if v.Field(i).CanInterface() { //是否需判断零值
			fieldName := v.Type().Field(i).Name
			fvalue := fmt.Sprintf("%v", v.Field(i).Interface())
			ftype := v.Type().Field(i).Tag.Get("search_type")
			fieldMap[fieldName] = objs.FieldInfo{ftype, fvalue}
		}
	}
	return fieldMap, nil
}
