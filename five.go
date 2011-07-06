// make a file from cmd
package main

import (
	"fmt"
	"os"
)

var (
	filename = "test2.txt"
)

func main() {
	f, err := os.Create(filename) // O_RDWR|O_CREATE|O_TRUNC, 0666
	// get used to writing error processing code
	if err != nil {
		fmt.Println("error : os.Create : " + err.String())
		os.Exit(1)
	}
	buf := make([]byte, 80) // reserve 80 bytes buffer space
	os.Stdin.Read(buf) // cmd window opens, waiting for input, <enter>
	f.Write(buf)
	if err != nil {
		fmt.Println("error : os.WriteString : " + err.String())
		os.Exit(1)
	}
	f.Close()
}

// the next step is to have cmd open for input of filename to create
