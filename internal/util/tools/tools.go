package tools

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/spf13/viper"
)

func DecodeConfig(configPath string, config any) error {
	fileName := filepath.Base(configPath)
	fileNames := strings.Split(fileName, ".")
	if len(fileNames) != 2 {
		return errors.New("fileNames len not equal 2")
	}

	vp := viper.New()
	vp.AddConfigPath(filepath.Dir(configPath))
	vp.SetConfigName(fileNames[0])
	vp.SetConfigType(fileNames[1])
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}
	err = vp.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}

func TimeCost() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}

func Str2Bytes(s string) []byte {
	if s == "" {
		return []byte{}
	}
	return (*[0x7fff0000]byte)(unsafe.Pointer(
		(*reflect.StringHeader)(unsafe.Pointer(&s)).Data),
	)[:len(s):len(s)]
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
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

type FieldInfo struct {
	Type  string
	Value string
}

func ConvStruct2Map(s interface{}) (map[string]FieldInfo, error) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("kind is not struct")
	}

	fieldMap := make(map[string]FieldInfo, 0)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct && v.Type().Field(i).Type.Name() == v.Type().Field(i).Name {
			if subFieldMap, err := ConvStruct2Map(v.Field(i).Interface()); err == nil {
				for fieldName, fieldInfo := range subFieldMap {
					fieldMap[fieldName] = fieldInfo
				}
			}
		} else if v.Field(i).CanInterface() { //TODO:是否需判断零值
			fieldName := v.Type().Field(i).Name
			fieldValue := fmt.Sprintf("%v", v.Field(i).Interface())
			fieldType := v.Type().Field(i).Tag.Get("search_type")
			fieldMap[fieldName] = FieldInfo{Type: fieldType, Value: fieldValue}
		}
	}
	return fieldMap, nil
}
