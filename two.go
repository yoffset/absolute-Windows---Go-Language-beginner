// write a string into file
package main

import (
	"os"
	"fmt"
)

var (
	filename = "test2.txt"
	content  = "test string"
)

// os import gives access to this func
// func Create(name string) (file *File, err Error)
// make local reference arg(ument) file, f for short

func main() {

	f, err := os.Create(filename) // O_RDWR|O_CREATE|O_TRUNC, 0666
	// get used to writing error processing code
	if err != nil {
		fmt.Println("error : os.Create : " + err.String())
		os.Exit(1)
	}
	// use f to write some content in the file with this method of File
	// func (file *File) WriteString(s string) (ret int, err Error)
	f.WriteString(content)
	if err != nil {
		fmt.Println("error : os.WriteString : " + err.String())
		os.Exit(1)
	}
	println(content)
	println(f)
	fmt.Println(f)
}

// notice all the times you've ran this, there is only one test string
// and the date for the file keeps advancing - O_TRUNC recreates clean if already exists
