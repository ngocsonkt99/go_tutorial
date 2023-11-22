package main

import "fmt"

func main() {
	//%v thay thế giá trị biến vào trong chuỗi, áp dụng cho Printf
	//if else
	number := 9

	//if condition { //code here }
	if number == 10 {
		fmt.Println("So bang:", number)
	} else {
		fmt.Println("So #", number)
	}

	if number < 10 {
		fmt.Printf("So %v be hon 10", number)
	} else {
		fmt.Printf("So bang %v", number)
	}
	fmt.Println("")

	println("===============================================================")
	//if statemen; condition { // code } gán và so sánh điều kiện luôn
	if a := 100; a > 100 {
		fmt.Println("a>100")
	} else {
		fmt.Println("a=100")
	}
	println("===============================================================")
	//Switch - Case
	So := 13
	switch So {
	case 1, 11, 111:
		fmt.Println("co so bat dau bang 1")
	case 2:
		fmt.Println("so bang 2")
	case 5:
		fmt.Println("so bang 3")
	case 13:
		fmt.Println("so bang 13")
	default:
		fmt.Println("unknown")
	}

	println("===============================================================")
	SoNumber := 19
	switch SoNumber {
	case 19:
		if SoNumber == 19 {
			goto handleNumberEqual10
		}
		fmt.Println("So bang 10")
	handleNumberEqual10:
		fmt.Println("Handle for case 9")

	case 9:
		fmt.Println("So 1")
		fallthrough
		//fallthrough chạy tiếp các case đồng thời nhận case sau chứ không break

	case 20:
		fmt.Println("case 20")
		fallthrough
	case 23:
		fmt.Println("case new")
	}

	println("===============================================================")

	//Loop
	//for ini; condition; post
	for i := 0; i < 10; i++ {
		if i == 4 {
			//break out vòng lặp
			//break
			//continue tiếp tục vòng lặp sau khi tới 4 và không print 4
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("Out Loop")

	println("===============================================================")

	//while
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	println("===============================================================")
	//infinite loop - lặp vô tận

	//for {
	//fmt.Println("infinite loop")
	//}
	println("===============================================================")
	//multiple for init, condition, post
	for i, j := 1, 2; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Println(i)
		fmt.Println(j)
	}
	println("===============================================================")
}
