// misc
package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	ps := os.Getpagesize()
	fmt.Println("page size = ", ps)

	sec, nsec, _ := os.Time()
	fmt.Println("time from 1 January 1970 =", sec, " secs.", nsec)

	t := time.LocalTime()
	fmt.Println(t)
}

// just a few things picked from
// http://golang.org/pkg/
// just to show how to get them
// most everyone should be away now
