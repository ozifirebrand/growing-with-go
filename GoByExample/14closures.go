package main

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println("new ints", newInts())
}

// anonymous function
func intSeq() func() int {
	i := 8
	return func() int {
		i++
		return i
	}
}
