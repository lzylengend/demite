package my_error

func WxError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "微信服务器错误", ErrorDesc: errorDesc}
}

func NotBindError() *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "没有绑定设备", ErrorDesc: ""}
}

func ExistApplyError() *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "已经提交申请", ErrorDesc: ""}
}
