package main

import "fmt"

//this "..." defines the variadic function
func sum(nums ...int) {
	fmt.Println("nums:",nums, " ")
	total := 0

	for _, num := range nums {
        total += num
    }
	fmt.Println("total:", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	numbers := []int{1, 2, 3, 4}
	sum(numbers...)
	sum(numbers...)

}
