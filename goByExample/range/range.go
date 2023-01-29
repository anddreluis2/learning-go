package main
import (
	"fmt"
)

func main() {
	nums := []int {2, 3, 4}
	sum := 0
	
	//take the sum of nums inside nums
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	//takes the index of num
	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	//take the first index and that value
	kvs := map[string]string{"a": "apple", "b" : "banana"}
	for k, v := range kvs {
		fmt.Println(k, "->", v)
	}

	for k := range kvs {
		fmt.Println("key", k)
	}

	for v := range kvs {
		fmt.Println("value", v)
	}

	  for i, c := range "go" {
        fmt.Println(i, c)
    }

}