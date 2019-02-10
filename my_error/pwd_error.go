package my_error

func PwdIdError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "数据库操作失败，id获取失败", ErrorDesc: errorDesc}
}

func PwdFailError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "密码错误", ErrorDesc: errorDesc}
}
