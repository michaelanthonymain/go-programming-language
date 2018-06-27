package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456789"))
}

func comma(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		buf.WriteString(string(v))
		if i != len(s)-1 && (len(s)-i-1)%3 == 0 {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}
