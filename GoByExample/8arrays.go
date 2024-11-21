package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a =", a)

	a[4] = 8
	fmt.Println("set:", a)
	fmt.Println("get index 4:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 3, 5, 1, 4}
	fmt.Println("b:", b)

	c := [...]int{1, 2, 4, 5, 9}
	fmt.Println("count:", c)

	d := [...]int{1, 5, 8, 3: 400, 89, 500}
	fmt.Println("idx:", d)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two D:", twoD)

	tooD := [2][3]int{
		{1, 2, 3},
		{2, 4, 6},
	}

	fmt.Println("2 D:", tooD)

}
