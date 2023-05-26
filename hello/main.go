package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
)

func main() {

	http.HandleFunc("/user/login", userLogin)
	// 主页面的地址 取这个目录下的index.html 作为主目录  但是如果在浏览器里输入  http://localhost:8080/main.go  会暴露你的文件
	//http.Handle("/", http.FileServer(http.Dir(".")))
	// 提供指定目录    http://localhost:8080/asset/
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/view/user", http.FileServer(http.Dir(".")))

	// http://localhost:8080/user/login.shtml
	http.HandleFunc("/user/login.shtml", func(writer http.ResponseWriter, request *http.Request) {
		// 请求这个url的时候返回本地的默认资源   view/user 前面不用加/  不然也是有问题的
		tpl, err := template.ParseFiles("view/user/login.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		// 这里的name  和 {{define xxxxx}} 里的名字相同
		tpl.ExecuteTemplate(writer, "/user/login.shtml", nil)
	})

	// 启动服务器 这里要加冒号 :8080
	http.ListenAndServe(":8080", nil)

}

func userLogin(writer http.ResponseWriter, request *http.Request) {
	//io.WriteString(writer, "hello word")

	// 解析参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	pwd := request.PostForm.Get("passwd")

	loginOK := false
	fmt.Println(reflect.TypeOf(mobile))
	fmt.Println("账号", mobile)
	fmt.Println("密码", pwd)

	if mobile == "18600000000" && pwd == "123456" {
		loginOK = true
	}
	str := "{\"code\":0,\"data\":{\"id\":1,\"token\":\"test\"}}"
	/// 这个方法是定义返回的data结构体
	data := make(map[string]interface{})
	data["id"] = 1
	data["gender"] = 0
	data["birthday"] = "2008-12-11"

	if !loginOK {
		str = "{\"code\":-1,\"msg\":\"密码不正确\"}}"
	}

	fmt.Println(str)
	Resp(writer, 0, data, "It's message")
}

type H struct {
	Code int         "json:\"code\""
	Msg  string      "json:\"msg\""
	Data interface{} "json:\"data\", omitempty"
}

func Resp(w http.ResponseWriter, code int, data interface{}, message string) {
	// 返回失败的json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Data: data,
		Msg:  message,
		Code: code,
	}

	ret, err := json.Marshal(h)
	if err != nil {

	}
	// 返回成功的json
	w.Write([]byte(ret))
}
