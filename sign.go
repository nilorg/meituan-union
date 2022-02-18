package union

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/url"
	"reflect"
	"sort"

	"github.com/nilorg/sdk/convert"
)

func signBuffer(params url.Values, signatureKey string) *bytes.Buffer {
	// 获取Key
	keys := []string{}
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	value := new(bytes.Buffer)
	value.WriteString(signatureKey)
	for _, k := range keys {
		value.WriteString(k)
		value.WriteString(params.Get(k))
	}
	value.WriteString(signatureKey)
	return value
}

// SignMD5 生成signMD5
func SignMD5(params url.Values, signatureKey string) string {
	value := signBuffer(params, signatureKey)
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, value)
	return hex.EncodeToString(h.Sum(nil))
}

func interfaceToString(src interface{}) string {
	if src == nil {
		panic(ErrTypeIsNil)
	}
	switch v := src.(type) {
	case string:
		return v
	case uint8, uint16, uint32, uint64, int, int8, int32, int64, float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// SignStructToParameter ...
func SignStructToParameter(value interface{}) (values url.Values, err error) {
	values = url.Values{}
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			xname := f.Tag.Get("json")
			if xname == "-" {
				continue
			}
			xvalue := interfaceToString(v.FieldByName(f.Name).Interface())
			if xvalue != "" {
				values.Set(xname, xvalue)
			}
		}
	case reflect.Ptr:
		for i := 0; i < t.Elem().NumField(); i++ {
			f := t.Elem().Field(i)
			xname := f.Tag.Get("json")
			if xname == "-" {
				continue
			}
			xvalue := interfaceToString(v.Elem().FieldByName(f.Name).Interface())
			if xvalue != "" {
				values.Set(xname, xvalue)
			}
		}
	default:
		err = ErrNotEqualStruct
	}
	return
}
