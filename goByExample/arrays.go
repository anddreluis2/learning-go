package main
import (
	"fmt"
)

func main() {
	var a [5] int
	fmt.Println("a = ", a)

	a[4] = 13
	fmt.Println("a = ", a)
	fmt.Println("len = ", len(a))

	b:= [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	var twoD[2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD = ", twoD)
}