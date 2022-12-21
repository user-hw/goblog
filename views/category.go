package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category

	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("路径不匹配，无法识别路径"))
		return
	}

	// 分页
	if err := r.ParseForm(); err != nil {
		categoryTemplate.WriteError(w, errors.New("xitongcuowu!!!"))
		return
	}
	pageString := r.Form.Get("page")
	if pageString == "" {
		pageString = "1"
	}
	page, _ := strconv.Atoi(pageString)
	// 每页显示的数量
	pageSize := 10

	categoryResponse, err := service.GetPostsByCategolyId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("出现错误"))
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
