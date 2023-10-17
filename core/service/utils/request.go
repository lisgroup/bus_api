package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewUtilsRequest(param RequestParam) *Request {
	return &Request{
		RequestParam: param,
		client:       &http.Client{},
	}
}

// RequestParam 请求工具类对象
type RequestParam struct {
	Url        string
	Body       interface{}
	Params     map[string]string
	Header     http.Header
	CaCertPath string
	CertFile   string
	KeyFile    string
}

type Request struct {
	client *http.Client
	RequestParam
	Response *http.Response
	Request  *http.Request
	Err      Errno
}

func (r *Request) Get() {
	uri, err := url.Parse(r.Url)
	if err != nil {
		r.Err = ErrReqParse.WithErr(err)
		AddLogger(r.Err, zap.String("url", r.Url))
		return
	}
	if len(r.Params) > 0 {
		urlRow := url.Values{}
		for key, value := range r.Params {
			urlRow.Add(key, value)
		}
		uri.RawQuery = urlRow.Encode()
	}

	r.Request, err = http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		r.Err = ErrReqNew.WithErr(err)
		AddLogger(r.Err, zap.String("uri", uri.String()))
		return
	}

	r.Request.Header = r.Header
	r.client.Timeout = time.Duration(time.Second * 10)
	r.Response, err = r.client.Do(r.Request)
	defer func() {
		r.client.CloseIdleConnections()
	}()
	if err != nil {
		r.Err = ErrHttpGet.WithErr(err)
		AddLogger(r.Err, zap.String("url", r.Url))
	}
}

func (r *Request) Post() {
	uri, err := url.Parse(r.Url)
	if err != nil {
		r.Err = ErrReqParse.WithErr(err)
		AddLogger(r.Err, zap.String("url", r.Url))
		return
	}
	if len(r.Params) > 0 {
		urlRow := url.Values{}
		for key, value := range r.Params {
			urlRow.Add(key, value)
		}
		uri.RawQuery = urlRow.Encode()
	}

	var byteList []byte
	byteList, err = json.Marshal(r.Body)
	if err != nil {
		r.Err = ErrJsonMarshal.WithErr(err)
		AddLogger(r.Err, zap.Field{Key: "body", Interface: r.Body, Type: zapcore.ObjectMarshalerType})
		return
	}
	r.Request, err = http.NewRequest(http.MethodPost, uri.String(), bytes.NewReader(byteList))
	if err != nil {
		r.Err = ErrReqNew.WithErr(err)
		AddLogger(r.Err, zap.String("uri", uri.String()))
		return
	}

	r.Request.Header = r.Header
	r.client.Timeout = time.Duration(time.Second * 10)
	r.Response, err = r.client.Do(r.Request)
	defer func() {
		r.client.CloseIdleConnections()
	}()
	if err != nil {
		r.Err = ErrHttpPost.WithErr(err)
		AddLogger(r.Err, zap.String("url", r.Url))
	}
}

func AddLogger(l Errno, fields ...zap.Field) {
	ctx := context.Background()
	switch l.Level {
	case zapcore.DebugLevel:
		logc.Debug(ctx, fields)
	case zapcore.ErrorLevel:
		fields = append(fields, zap.Stack("stack"))
		logc.Error(ctx, fields)
	default:
		logc.Info(ctx, fields)
	}

}
