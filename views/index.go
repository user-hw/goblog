package views

import (
	"errors"

	"goblog/common"
	"goblog/service"
	"log"
	"net/http"
	"strconv"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	if err := r.ParseForm(); err != nil {
		index.WriteError(w, errors.New("xitongcuowu!!!"))
		return
	}
	pageString := r.Form.Get("page")
	page := 1
	if pageString != "" {
		page, _ = strconv.Atoi(pageString)
	}
	// 每页显示的数量
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	index.WriteData(w, hr)

}
