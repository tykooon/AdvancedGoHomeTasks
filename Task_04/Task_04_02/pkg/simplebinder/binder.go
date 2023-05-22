package simplebinder

import (
	"reflect"
	"strconv"
)

func Bind(data map[string]string, obj any) {
	if len(data) == 0 || obj == nil {
		return
	}
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	if objType.Kind() != reflect.Pointer || objVal.Elem().Kind() != reflect.Struct {
		return
	}
	// ----------- Only pointers on structures in further code

	for i := 0; i < objVal.Elem().NumField(); i++ {
		fieldData := objType.Elem().Field(i)
		fieldValue := objVal.Elem().Field(i)
		var fieldStr string
		tag, ok := fieldData.Tag.Lookup("name")
		if ok {
			fieldStr, ok = data[tag]
			if ok && fieldStr != "" {
				SetFieldValue(&fieldValue, fieldData, fieldStr)
			}
		} else if fieldStr, ok := data[fieldData.Name]; ok {
			SetFieldValue(&fieldValue, fieldData, fieldStr)
		}
	}
}

func SetFieldValue(fieldValue *reflect.Value, fieldData reflect.StructField, value string) {
	if !fieldValue.CanSet() {
		return
	}
	switch fieldData.Type.Kind() {
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Int:
		intValue, err := strconv.ParseInt(value, 10, 16)
		if err == nil {
			fieldValue.SetInt(intValue)
		}
	case reflect.Uint64:
	case reflect.Uint32:
	case reflect.Uint16:
	case reflect.Uint:
		uintValue, err := strconv.ParseUint(value, 10, 16)
		if err == nil {
			fieldValue.SetUint(uintValue)
		}
	case reflect.Float32:
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err == nil {
			fieldValue.SetFloat(floatValue)
		}
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			fieldValue.SetBool(boolValue)
		}
	default:
		return
	}
}
