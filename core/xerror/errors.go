package xerror

/**
常用通用固定错误
*/

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// GetCode 返回给前端的错误码
func (e *CodeError) GetCode() int {
	return e.Code
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Msg
}

// CodeErrorResponse 此处结构体用于统一返回结果
type CodeErrorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

// NewDefaultError 设置默认错误码
func NewDefaultError(msg string) error {
	return NewCodeError(ServerCommonError, msg)
}

// NewParamsFailedError 设置接口参数json校验错误码数据
func NewParamsFailedError(msg string) error {
	return NewCodeError(RequestParamError, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

// Data 返回通用错误数据
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

// NewSuccessJson 接口请求成功返回数据
func NewSuccessJson(resp interface{}) *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: OK,
		Msg:  MapErrMsg(OK),
		Data: resp,
	}
}
