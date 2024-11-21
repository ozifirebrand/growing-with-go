package main

import "fmt"

func main() {
	const rangeValue = 3
	p := 1
	for p <= 3 {
		fmt.Println(p)
		p = p + 1
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	for p := range rangeValue {
		fmt.Println("range", p)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println()
	}

	for i := range 3 {
		fmt.Println("Range ", i)

	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for {
		fmt.Println("LOOP")
		break
	}
}
