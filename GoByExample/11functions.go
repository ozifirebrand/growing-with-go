package main

import (
	"fmt"
)

func main() {
	fmt.Println("Plus:", plus(4, 8))
	fmt.Println("Plus plus:", plusPlus(4, 8, -5))

}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a int, b int, c int) int {
	return a + b + c
}
