package main

import (
	"fmt"
)

func main() {
	binSum := 50 + 0b10
	octoSum := 50 + 0o10
	hexSum := 50 + 0x10

	fmt.Println(binSum)  // 52
	fmt.Println(octoSum) // 58
	fmt.Println(hexSum)  // 66
}
