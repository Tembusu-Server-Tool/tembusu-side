package main

import (
	"fmt"
	"os/exec"
	"net/http"
	// "time"
	// "bufio"
	// "strings"
)

func cors(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
        w.Header().Add("Access-Control-Allow-Credentials", "true") //设置为true，允许ajax异步请求带cookie信息
        w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
        w.Header().Set("content-type", "application/json;charset=UTF-8")             //返回数据格式是json
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        f(w, r)
    }
}

func handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	cmd := exec.Command("sinfo", "--Node", "--format=\"%8N %10P %5T %5c %8O\"")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("combined out:\n%s\n", out)

	fmt.Fprintf(w, string(out))
}

func main() {
	http.HandleFunc("/check", handle)
	err := http.ListenAndServe("192.168.51.112:4000", cors(handle))
	if err != nil {
		fmt.Println(err)
	}
}
