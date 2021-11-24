package union

import "testing"

func TestStructToParameter(t *testing.T) {
	value := struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"-"`
	}{
		A: "value a",
		B: "value b",
	}
	params, err := SignStructToParameter(value)
	if err != nil {
		t.Errorf("反射类型错误：%v", err)
		return
	}
	t.Log(params)
}
