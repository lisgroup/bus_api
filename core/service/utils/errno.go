package utils

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

type Errno struct {
	Code  int           `json:"code"`
	Msg   string        `json:"msg"`
	Err   error         `json:"error"`
	Level zapcore.Level `json:"level"`
}

func (e Errno) Error() string {
	return fmt.Sprintf("code:%d,msg:%s,err:%s", e.Code, e.Msg, e.Err)
}

func (e Errno) WithErr(err error) Errno {
	e.Err = err
	// e = errors.New("")
	return e
}

var (
	ErrJsonMarshal = Errno{10022, "序列化json失败", nil, zapcore.ErrorLevel}
	ErrHttpGet     = Errno{10101, "http get 请求失败", nil, zapcore.WarnLevel}
	ErrHttpPost    = Errno{10102, "http post 请求失败", nil, zapcore.WarnLevel}
	ErrReqParse    = Errno{10110, "url parse err", nil, zapcore.ErrorLevel}
	ErrReqNew      = Errno{10111, "new request err", nil, zapcore.ErrorLevel}
)
