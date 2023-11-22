package main

import "fmt"

func main() {
	//Khai bao single variable
	var number int
	number = 20
	fmt.Println(number)

	var number1 int = 40
	fmt.Println(number1)

	//Type inference
	var email = "rybuson@gmail.com"
	fmt.Println(email)

	println("===================================================")
	//Khai bao nhieu bien
	//Cung type
	var a, b int
	a = 1
	b = 3
	fmt.Println(a)
	fmt.Println(b)
	//variable
	var a1, b1 int = 10, 11
	fmt.Println(a1)
	fmt.Println(b1)

	var a2, b2 = 12, 14
	fmt.Println(a2)
	fmt.Println(b2)

	println("===================================================")
	//Khac type
	var (
		name    string
		address string
		age     int
	)

	name = "Rybu"
	address = "Australia"
	age = 23

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)
	println("====================")

	var (
		name1    string = "Son"
		address1 string = "VietNam"
		age1     int    = 24
	)

	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age1)
	println("====================")

	var name2, address2, age2 = "Ung", "15 Vo Van Kiet", 26
	fmt.Println(name2)
	fmt.Println(address2)
	fmt.Println(age2)
	println("====================")

	language := "Golang"
	fmt.Println(language)
}
