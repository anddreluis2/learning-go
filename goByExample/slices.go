package main
import (
	"fmt"
)

func main() {
	//first argument: datatype, second argument: array length
	s:= make([]string, 3)
	fmt.Println(s)

	s[0]= "a"
	s[1]= "b"
	s[2]= "c"

	fmt.Println(s)
	fmt.Println(s[2])

	s= append(s, "d")
	s= append(s, "e", "f")
	fmt.Println(s)


	c:= make([]string, len(s))
	//copy the last state
	copy(c, s)
	fmt.Println(c)

	//slices the interval between the index
	l := s[2:5]
	fmt.Println(l)

	//slices everything after the index
	l = s[:5]
	fmt.Println("dois ponto cinco",l)


	//slices everything before the index
	l = s[2:]
	fmt.Println("dois dois ponto",l)
	
	//slices the index
	l = s[:5]
	
		t := make([][]int, 3)
		for i := 0; i <3; i++ {
			innerLen := i+1
			t[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				t[i][j] = i + j
			}
		}
			fmt.Println("t: ", t)
}