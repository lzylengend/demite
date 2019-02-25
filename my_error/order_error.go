package my_error

func GoodNotEnoughError(productName string) *ErrorCommon {
	return &ErrorCommon{HasError: false, ErrorShowDesc: productName + " 商品存货不够", ErrorDesc: ""}
}
