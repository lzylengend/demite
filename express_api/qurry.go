package express_api

import (
	"demite/util"
	"fmt"
)

func QurryExpress(expressId string) {
	body, err := util.ClientDo("GET", "http://q.kdpt.net/api?id=testkey&com=zhongtong&nu=370817399305", []byte{})
	if err != nil {
		return
	}

	fmt.Println(string(body))
}
