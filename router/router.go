package router

import (
	"goblog/api"
	"goblog/views"
	"html/template"
	"net/http"
	"os"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	t := template.New("admin.html")
	path, _ := os.Getwd()
	// fmt.Printf("path: %v\n", path)
	t, _ = t.ParseFiles(path + "/template/admin.html")

	t.Execute(w, "")
}

func Router() {
	//1.页面 views 2.数据 api (json) 3.静态资源

	http.HandleFunc("/", views.HTML.Index)

	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login/", views.HTML.Login)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing/", views.HTML.Writing)
	http.HandleFunc("/pigeonhole/", views.HTML.Pigeonhole)
	// http.HandleFunc("/userinfo/", views.HTML.UserInfo)

	http.HandleFunc("/admin/", Admin)

	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
