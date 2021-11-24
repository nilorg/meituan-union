package union

const (
	// BaseURL ...
	BaseURL = "https://openapi.meituan.com/api"
)

type ResponseReturn struct {
	ErrCode string `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
