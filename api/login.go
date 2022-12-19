package api

import (
	"goblog/common"
	"goblog/service"
	"log"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接收用户名和密码 返回 对应的json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	LoginRes, err := service.Login(userName, passwd)
	log.Println(LoginRes)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, LoginRes)
}
