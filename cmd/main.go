package main

import (
	"fmt"
	"hello" // Updated the go.mod init file from module letGO to hello
	"os"
)

// At the root, go run ./cmd <Arg>
func main() {
	if len(os.Args) > 1 {
		// os.Args[0] is the programme name
		fmt.Println(hello.Say(os.Args[1]))
	} else {
		fmt.Println(hello.Say("World!"))
	}

}
