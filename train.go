package main

import "fmt"

func main() {
	var a, z int
	count := 0
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&z)
		if z == 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
