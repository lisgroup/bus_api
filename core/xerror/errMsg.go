package xerror

var message map[int]string

func init() {
	message = make(map[int]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode int) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
