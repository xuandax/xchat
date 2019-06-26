package main

import (
	"github.com/xuanxiaox/xchat/controller"
	"html/template"
	"net/http"
)

//注册所有的模板
func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		panic(err)
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			err := v.ExecuteTemplate(writer, tplName, nil)
			if err != nil {
				panic(err)
			}
		})
	}
}

func main() {

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)

	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	RegisterView()
	http.ListenAndServe(":8080", nil)
}
