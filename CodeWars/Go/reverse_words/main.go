package main

import (
	"fmt"
)

func main() {
	ret := ReverseWords("The  quick brown fox jumps over the lazy dog. ")
	if ret != "ehT  kciuq nworb xof spmuj revo eht yzal .god " {
		fmt.Printf("Failed: %s", ret)
	}
}

func ReverseWords(str string) string {
	var word string
	var rev string

	for _, e := range str {
		if e != ' ' {
			word = string(e) + word
		} else {
			rev = rev + word + " "
			word = ""
		}
	}
	return rev + word
}
