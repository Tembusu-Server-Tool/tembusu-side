package main

import (
	"fmt"
	// "time"
	// "bufio"
	"strings"
)

func parse(out string) {
	lines := strings.Split(out, "\n")
	lines = lines[0:len(lines) - 1]
	for index, _ := range lines {
		lines[index] = lines[index][1:len(lines[index]) - 1]
	}

	for _, line := range lines {
		elements := strings.Split(strings.Join(strings.Fields(line), " "), " ")
		fmt.Println(elements)

	}

}


func main() {
	abc := "\"xcne   c   abc   d\"\n\"xcnd   d idle ccc\"\n"
	parse(abc)
}

