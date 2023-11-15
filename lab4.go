package main

import "fmt"

// func func_name(params) return type { //code here }
func chao() {
	fmt.Println("Hello guy")

}

func Chao(name string) {
	fmt.Println("Hello", name)
}

func greeting(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}

// Multiple return values
func rectInfo(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

func rectInfo2(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main() {
	chao()
	Chao("Rybu")
	result := greeting("Son")
	fmt.Println(result)

	println("===================================================")

	w, h, area := rectInfo(100, 200)
	fmt.Println("width=", w)
	fmt.Println("height=", h)
	fmt.Println("area=", area)

	println("===================================================")

	w, h, isSquare := rectInfo2(300, 300)
	if isSquare {
		fmt.Println("This is Square")

	} else {
		fmt.Println("width=", w)
		fmt.Println("height=", h)
	}
}
