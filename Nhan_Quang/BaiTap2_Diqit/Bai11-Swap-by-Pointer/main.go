// mã giả hàm swap(a *int, b *int)
/*
Hàm swap(a, b)
    temp = giá trị tại a
    giá trị tại a = giá trị tại b
    giá trị tại b = temp
Kết thúc hàm
*/

package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	fmt.Println("Trước khi swap:")
	fmt.Println("x =", x, ", y =", y)

	swap(&x, &y)

	fmt.Println("Sau khi swap:")
	fmt.Println("x =", x, ", y =", y)
}
