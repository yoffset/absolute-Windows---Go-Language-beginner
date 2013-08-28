// fool around a bit with data in file
package main

import (
	"os"
	"fmt"
)

var (
	filename = "test2.txt"
	content  = "test string \n"             // notice format char
	contentB = []byte("1234567890    5  8") // Write expects bytes
	contentC = "back up"
)


func main() {

	f, err := os.Create(filename) // O_RDWR|O_CREATE|O_TRUNC, 0666
	// get used to writing error processing code
	if err != nil {
		fmt.Println("error : os.Create : " + err.String())
		os.Exit(1)
	}

	_, err := f.WriteString(content)
	if err != nil {
		fmt.Println("error : os.WriteString : " + err.String())
		os.Exit(1)
	}

	_, err := f.Write(contentB) // notice 8 on end
	if err != nil {
		fmt.Println("error : os.Write : " + err.String())
		os.Exit(1)
	}

	f.Seek(-2, 2)           // backs up 2 bytes for next write start
	_, err := f.WriteString(contentC) // backed up over 8 - how about that?
	if err != nil {
		fmt.Println("error : os.WriteString : " + err.String())
		os.Exit(1)
	}
	f.Close() // close file - release memory
}
