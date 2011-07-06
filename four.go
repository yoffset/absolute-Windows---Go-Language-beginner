// read mystery file
package main

import (
	"os"
	"fmt"
	"io" // new lib
)

var (
	madeFile = "test.txt" // mystery file
)


func main() {
	BufIsFileSize(madeFile) // makes buffer size of file contents
}



func BufIsFileSize(filename string) ([]byte, os.Error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()    // when return is next, defer jumps que
	d, err := f.Stat() // file statistics
	if err != nil {
		return nil, err
	}
	buf := make([]byte, d.Size)  // buf size of file statistic size
	_, err = io.ReadFull(f, buf) // read all file, of course, into buf
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n%s\n", buf) // file contents - try it without %s
	return buf, err
}
