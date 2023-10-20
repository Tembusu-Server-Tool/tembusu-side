package main

import (
	"fmt"
	"os/exec"
	"net/http"
	// "time"
	// "bufio"
	"strings"
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

func handleCheck(w http.ResponseWriter, r *http.Request) {
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

type machineInfo struct {
	name string
	status string
}
// Parse string to [{machineId, }] and select all long and idle machine
func parse(out string) {
	lines := strings.Split(out, "\n")
	for index, _ := range lines {
		lines[index] = lines[index][1:-1]
	}
	output := make()
	for _, line := range lines {
		elements := strings.Split(line, " ")
		
	}

}

func handlePredict(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	cmd := exec.Command("sinfo", "--Node", "--format=\"%8N %10P %5T %5c %8O\"")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	parsedOut = parse(out)
}

func main() {
	http.HandleFunc("/check", handleCheck)
	http.HandleFunc("/predict", handlePredict)
	http.ListenAndServe("192.168.51.112:4000", cors(handleCheck))
	http.ListenAndServe("192.168.51.112:4000", cors(handlePredict))
}
