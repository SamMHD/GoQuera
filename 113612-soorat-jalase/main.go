package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// func to determine if a character is a number
func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {
	var suffixes = []string{"th", "st", "nd", "rd", "th", "th", "th", "th", "th", "th"}

	// read from stdin
	reader := bufio.NewReader(os.Stdin)
	for text, err := reader.ReadString('\n'); err == nil || err == io.EOF; text, err = reader.ReadString('\n') {
		res := ""
		for i, c := range text {
			res += string(c)
			if isNumber(text[i]) && ((i+1 < len(text) && !isNumber(text[i+1])) || i+1 == len(text)) {
				if i > 0 && text[i-1] == '1' {
					res += "th"
				} else {
					res += suffixes[text[i]-'0']
				}
			}
		}
		fmt.Print(res)
		if err == io.EOF {
			fmt.Print("\n")
			return
		}
	}
}
