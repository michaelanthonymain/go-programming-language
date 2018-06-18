package echo

import (
	"fmt"
	"os"
	"strings"
)

// Echo3 ...
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
