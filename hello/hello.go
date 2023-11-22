package main

import "fmt"

func main() {
	fmt.Println("Hello Golang!")

	fmt.Println(Sum(5))

	var name string
	name = "Ung Ngoc Son"
	fmt.Println("Hi", name)

	// cach #
	name1 := "Son Ung"
	var name2 string = "Son"
	var name3 = "Ngoc Son"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)

	name4, age := "Mai", 23
	fmt.Println(name4, age)
}

func Sum(number int) int {
	sum := 0
	for i := 0; i <= number; i++ {
		sum += i
	}
	return sum
}
