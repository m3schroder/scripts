package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var directory string

	if os.Args != nil && len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		directory = "./"
	}

	fmt.Printf("Number of lines in %s\n", directory)
	app := "rg"

	arg0 := "-c"
	arg1 := ""
	arg2 := directory

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	responses := strings.Split(string(stdout[:]), "\n")

	count := 0
	for _, r := range responses {
		split := strings.Split(r, ":")
		rowCount, _ := strconv.Atoi(split[len(split)-1])
		count += rowCount
	}
	fmt.Println(count)
}
