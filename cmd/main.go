package main

import (
	"fmt"
	"hello" // Updated the go.mod init file from module letGO to hello
	"os"
)

// At the root, go run ./cmd <Arg>
func main() {
	fmt.Println(hello.Say(os.Args[1:]))

}
