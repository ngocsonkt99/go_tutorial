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

	println("===================================================")

	for i, j := 0, 5; i <= 5 && j >= 0; i, j = i+1, j-1 {
		fmt.Printf("%d %d\n", i, j)
	}

	println("===================================================")

	// khai bao don gian
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	// khai bao short hand
	b := [3]int{1, 2, 3}
	// bo qua chieu dai cua mang
	c := [...]int{6, 4, 34}

	for i, v := range b {
		fmt.Printf("Phan tu thu %d la: %d\n", i, v)
	}
	println("===================================================")
	c[0] = 66
	fmt.Println("Mang C", b)
	println("===================================================")

	d := [5]int{1,3,6,8,9}
	var e []int = d[1:4] // slice out [3 6 8]
	fmt.Println(e)

	arr := [...]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // [2, 3, 4]
	for i, _ := range slice {
		slice[i]++
	}
	fmt.Println(slice) // [3, 4, 5]
	fmt.Println(arr) // [1, 3, 4, 5, 5]

	array := [...]int{2, 3, 4, 5, 6}
	slice2 := array[1:4] // [3 4 5]
	fmt.Println("Do dai cua slice2 la", len(slice2), "va suc chua la", cap(slice2))
	//do dai 3 vi slice dc tao tu index 1,2,3 cua mang arr, suc chua 4 vi slice bat dau tu index 1, mang array tu index 1 den het mang co 4 phan tu
	slice2 = array[0: cap(slice2)]
	//đoạn mã sẽ tạo một slice mới slice2 từ array, bắt đầu từ phần tử đầu tiên của array và kết thúc tại vị trí có chỉ số là cap(slice2) - 1
	//chúng ta chỉ định rằng slice mới slice2 sẽ có độ dài bằng sức chứa của slice2
	fmt.Println("Slice2 sau khi thay doi co do dai la:", len(slice), "va suc chua la:", cap(slice2))


	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice3 co do dai la", len(slice3), "va suc chua la", cap(slice3))
	// slice co do dai la 5 va suc chua la 5
	slice3 = append(slice3, 6) //su dung append se gap doi suc chua ban dau neu khong du. tất cả các phần tử của slice ban đầu sẽ được sao chép sang slice mới và phần tử 6 sẽ được thêm vào cuối slice mới. Kết quả là, slice mới chứa các phần tử từ 1 đến 6.
	// Sau khi append() hoàn thành, slice slice3 sẽ trỏ đến slice mới và slice ban đầu sẽ không còn được sử dụng.
	fmt.Println("slice3 sau khi thay doi co do dai la", len(slice3), "va suc chua la", cap(slice3))
	fmt.Println(slice3)
	// slice sau khi thay doi co do dai la 6 va suc chua la 10
	// Suc chua cua slice da duoc tang len gap doi
}
