package main

import "fmt"

func main() {
	variadicFunctionSum(1, 4, 6, 2)
	variadicFunctionSum(16, 20)
	nums := []int{1, 9, 0, 2}
	variadicFunctionSum(nums...)

}

func variadicFunctionSum(nums ...int) {
	fmt.Println("nums:", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)

}
