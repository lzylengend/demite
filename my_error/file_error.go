package my_error

func FileParseError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "文件解析错误", ErrorDesc: errorDesc}
}

func FileWriteError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "文件写入错误", ErrorDesc: errorDesc}
}

func FileReadError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "文件读取错误", ErrorDesc: errorDesc}
}
