package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
* Exercises:
* x) include the name of the command that invoked Args
* x) print the index of each argument
* 3) measure differences in running time between versions
**/

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		i := strconv.Itoa(index)
		s += sep + arg + sep + i
		sep = " "
	}
	fmt.Println(s)
}
