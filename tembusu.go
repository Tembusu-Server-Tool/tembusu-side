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
	var conn *net.TCPConn
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "172.26.191.78:4000")
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
	}

	reader := bufio.NewReader(conn)
	for {
		data, _ := reader.ReadString('\n')
		content := strings.Replace(string(data), "\n", "", -1)

		if (content == "request") {
			cmd := exec.Command("ls")
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("combined out:\n%s\n", string(out))
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
			fmt.Printf("combined out:\n%s\n", string(out))
			conn.Write([]byte(out))
		}
	}


}