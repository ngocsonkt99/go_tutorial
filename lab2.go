package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	//%t: Định dạng boolean. Khi sử dụng %t, giá trị boolean true sẽ được in ra là "true", và giá trị boolean false sẽ được in ra là "false".
	//
	//%c: Định dạng ký tự (rune). Khi sử dụng %c, giá trị rune sẽ được in ra dưới dạng ký tự Unicode tương ứng.
	//
	//%u: Định dạng số nguyên không dấu (unsigned integer). Khi sử dụng %u, giá trị số nguyên không dấu sẽ được in ra dưới dạng số nguyên không dấu.
	//
	//%d: Định dạng số nguyên có dấu (signed integer). Khi sử dụng %d, giá trị số nguyên có dấu sẽ được in ra dưới dạng số nguyên có dấu.

	//Boolean logic
	var myBool bool = true // false
	fmt.Print(myBool)

	var mySecondBool bool = false
	fmt.Println(mySecondBool)

	//String chuỗi
	var myString string = "hi"
	fmt.Println(myString)

	//Int số nguyên
	var myInt int = 23
	fmt.Println(myInt)

	//Range
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	//Result Range
	//-128
	//127

	//-32768
	//32767

	//-2147483648
	//2147483647

	//-9223372036854775808
	//9223372036854775807

	//Bits
	fmt.Println("============")
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println(bits.OnesCount16(math.MaxUint16))
	fmt.Println(bits.OnesCount32(math.MaxUint32))
	fmt.Println(bits.OnesCount64(math.MaxUint64))

	//Uint
	fmt.Println("==============")
	var myUint uint = 10
	fmt.Println(myUint)

	fmt.Println("==============")
	var myByte byte = 1
	fmt.Println(myByte)
	fmt.Printf("%T", myByte)
	//fmt.Println(math.MaxUint8)

	fmt.Println("==============")
	var a byte = 'E'
	fmt.Printf("%X", a)
	//Trong hệ thập lục phân, giá trị của ký tự 'E' là 45
	//%X được sử dụng để in ra giá trị số nguyên không dấu (unsigned integer) dưới dạng chuỗi thập lục phân

	//Float - Số thực
	fmt.Println("==============")
	var myFloat float64 = 10.02
	fmt.Println(myFloat)

	//complex z = a +bi
	var z1 complex64 = 10 + 1i
	fmt.Println(z1)

	var z2 complex64 = 10 + 1i
	fmt.Println(z1 + z2)

	//Rune - alias for Int32
	var myString2 string = "Sơn"
	fmt.Println(myString2)
	//Dùng runes chứ không đưa thẳng vào để đọc theo format %c định dạng ký tự rune
	runes := []rune(myString2)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	fmt.Println("==============")
	var myRune rune = 'A'
	fmt.Printf("%T", myRune)
	//%T trong Printf, ta in ra kiểu dữ liệu của biến myRune

	fmt.Println("==============")
	var mRune rune = '�'
	fmt.Printf("%c - %u - %d", mRune, mRune, mRune)
	//In ra ký tự '�' (ký tự không hợp lệ trong Unicode), giá trị Unicode của ký tự đó là 65533 và giá trị số nguyên có dấu của ký tự đó cũng là 65533

	const PI = 3.14159
	//Nếu khai báo hằng số thì không thể gán lại
	//PI := 123
	fmt.Println(PI)
}
