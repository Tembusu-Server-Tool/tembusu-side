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
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "localhost:4000")
	var conn *net.TCPConn
	var err error
	for {
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if (err != nil) {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		} else {
			break
		}
	}
	fmt.Println("connected!")

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