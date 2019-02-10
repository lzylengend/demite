package my_error

func UserNameError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "用户名错误", ErrorDesc: errorDesc}
}

func UserNameExistError() *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "用户名已存在", ErrorDesc: ""}
}
