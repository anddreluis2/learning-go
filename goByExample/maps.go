package main
import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	//create a tuple with the values
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	//v1 assumes k2 value.
	v1 := m["k2"]
	fmt.Println(v1)

	fmt.Println("len:", len(m))

	//ok
	delete(m, "k1")
	fmt.Println(m)
}