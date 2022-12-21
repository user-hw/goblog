package views

import (
	"fmt"
	"goblog/common"

	"goblog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	fmt.Println("正在调用这个函数。。。。。。。")

	pigeonholeRes := service.FindPostPigeonhole

	pigeonhole.WriteData(w, pigeonholeRes)
}
