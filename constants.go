package union

const (
	OK      = "ok"
	OKCode  = "0"
	Err     = "err"
	ErrCode = "1"
)

var (
	CallbackOK  = &ResponseReturn{ErrCode: OKCode, ErrMsg: OK}
	CallbackErr = &ResponseReturn{ErrCode: ErrCode, ErrMsg: Err}
)
