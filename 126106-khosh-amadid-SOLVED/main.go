package main

import (
	"fmt"
)

func HelloCodeCup(n int) string {
	return fmt.Sprintln("Hello CodeCup", n)
}

func main() {
	str := HelloCodeCup(6)
	fmt.Println(str)
}
