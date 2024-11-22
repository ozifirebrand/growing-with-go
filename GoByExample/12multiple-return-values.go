package main

import "fmt"

func main() {
	c, d := multipleReturnValues(6, 2, -9)
	fmt.Println("Multiple return values", c, d)

}

func multipleReturnValues(a int, b int, c int) (int, int) {
	return a, b + c
}
