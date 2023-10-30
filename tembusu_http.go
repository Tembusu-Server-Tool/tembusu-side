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
// Parse string to [{machineId, }]
// Example: [xgph19 long mixed 64 2.10]
func parse(out string) [][]string {
	lines := strings.Split(out, "\n")
	for index, _ := range lines {
		lines[index] = lines[index][1:len(lines[index]) - 1]
	}
	parsedOut := make([][]string, len(lines))
	for index, line := range lines {
		elements := strings.Split(strings.Join(strings.Fields(line), " "), " ")
		parsedOut[index] = elements
	}
	return parsedOut
}

// Filter long and idle machines with most CPUs
func filter(machines [][]string) []string {
	fmt.Println(machines)
	var filteredMachines []string
	for _, machine := range machines {
		if machine[0] == "xcnc" {
			continue
		}
		if machine[0] == "xcnd" {
			continue
		}
		if machine[1] == "long"  {
			if machine[2] == "idle" {
				filteredMachines = append(filteredMachines, machine[0])
			}
		}
	}
	return filteredMachines

}

// Convert [{machines, }] to machinesId,machinesId which can be used in script
func generate(filteredMachines []string) string {
	size := len(filteredMachines)
	if (size < 20) {
		return "The number of long and idle machines is not enough"
	}
	result := ""
	for i := size - 1; i >= 0; i-- {
		result += filteredMachines[i]
	}
	return result
}

func handlePredict(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	cmd := exec.Command("sinfo", "--Node", "--format=\"%8N %10P %5T %5c %8O\"")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	parsedOut := parse(string(out))
	filteredResult := filter(parsedOut)
	finalResult := generate(filteredResult)

	fmt.Fprintf(w, finalResult)
}

func main() {
	out := "\"xgph7    long     idle 64    2.55\"\n\"xgph7    long       mixed 64    2.55 \""
	parsedOut := parse(out)
	fmt.Println(parsedOut)
	filteredResult := filter(parsedOut)
	fmt.Println(filteredResult)
	// finalResult := generate(filteredResult)
	// http.HandleFunc("/check", handleCheck)
	// http.HandleFunc("/predict", handlePredict)
	// http.ListenAndServe("192.168.51.112:4000", cors(handleCheck))
	// http.ListenAndServe("192.168.51.112:4000", cors(handlePredict))
}
