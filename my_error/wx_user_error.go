package my_error

func WxError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "微信服务器错误", ErrorDesc: errorDesc}
}
