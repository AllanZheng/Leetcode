package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {

	command = strings.Replace(command, "(al)", "al", -1)
	command = strings.Replace(command, "()", "o", -1)
	return command
}

func main() {
	fmt.Println(interpret("(al)G(al)()()G"))
}
