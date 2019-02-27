package my_error

func GoodCodeExistError() *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "唯一码重复", ErrorDesc: ""}
}
