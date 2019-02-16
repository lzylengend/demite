package my_error

type ErrorCommon struct {
	ErrorDesc     string `json:"errordesc"`
	ErrorShowDesc string `json:"errorshowdesc"`
	HasError      bool   `json:"haserror"`
}

func NoError() *ErrorCommon {
	return &ErrorCommon{HasError: false, ErrorShowDesc: "", ErrorDesc: ""}
}

func JsonError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "上传数据错误", ErrorDesc: errorDesc}
}

func NoLoginError() *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "未登录", ErrorDesc: ""}
}

func DbError(errorDesc string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "数据库错误", ErrorDesc: errorDesc}
}

func NotNilError(key string) *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: key + "不能为空", ErrorDesc: ""}
}
