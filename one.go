// create a file
package main
// every prog needs package to identify relationship with other files, or not

// these are installed program packages we are relating to
import (
	"os"
	"fmt"
)

var (
	filename = "test2.txt"
)

// // is a comment line that compiler ignores

// os import gives access to this function
// func Create(name string) (file *File, err Error)
// *File is a pointer to internal type struct(ure) File 
// returns two args, one we'll toss away (_)for now


func main() {

	_, err := os.Create(filename) // O_RDWR|O_CREATE|O_TRUNC, 0666
	// get used to writing error processing code
	if err != nil {
		fmt.Println("error : os.Create : " + err.String())
		os.Exit(1)
	}
}

// press the light blue V triangle button to check if compile or error

// if compile, press exclamation mark

// see files created

// press one left of exclamation to clean up files