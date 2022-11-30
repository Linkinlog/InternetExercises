package main

import (
	"strconv"
	"strings"
)

func main() {
	//
}
func Is_valid_ip(ip string) bool {
	// Break string out into 4 slices
	str := strings.Split(ip, ".")
	// Make sure that the IP address has 4 octets
	if len(str) != 4 {
		return false
	}
	for _, e := range str {
		i, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		// Make sure that each octet is between 0-255
		if i > 255 || i < 0 {
			return false
		}
		// Make sure there are no leading 0s
	}
	return false
}
