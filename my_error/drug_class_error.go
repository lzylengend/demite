package my_error

func DrugClassDelError() *ErrorCommon {
	return &ErrorCommon{HasError: true, ErrorShowDesc: "分类下不为空，不可删除", ErrorDesc: ""}
}
