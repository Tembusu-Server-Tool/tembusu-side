package main

import (
	"fmt"
	"log"
	"os/exec"
	"net"
	"time"
	"bufio"
	"strings"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.51.112:4000")
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	} 
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
		go handle(tcpConn)
	}
}

func handle(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		data, _ := reader.ReadString('\n')
		content := strings.Replace(string(data), "\n", "", -1)

		if (content == "check") {
			cmd := exec.Command("sinfo", "--Node", "--format=\"%10N %.6D %10P %10T %20E %.4c %.8z %8O %.6m %10e %.6w %.60f\"")
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("combined out:\n%s\n", string(out))
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
			fmt.Printf("combined out:\n%s\n", []byte(string(out)))
			conn.Write([]byte(string(out)))
		}
	}
}