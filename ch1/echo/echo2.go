package echo

import (
	"fmt"
	"os"
	"strconv"
)

// Echo2 ...
func Echo2() {
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		i := strconv.Itoa(index)
		s += sep + arg + sep + i
		sep = " "
	}
	fmt.Println(s)
}
