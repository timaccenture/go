package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {

	// arrays
	var a [4]int
	fmt.Println("following array is empty:", a)

	b := [3]int{1, 2, 3}
	fmt.Println("following array is initialized:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d array looks as follows:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d array initliazed:", twoD)

	//slices (Python lists),
	//typed only by the elements they contain (not the number of elements as with arrays)

	var s []string
	fmt.Println("empty slice:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty slice:", s, "len:", len(s), "capacity:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d", "e")

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("original string:", s)
	fmt.Println("copied string:", c)
	fmt.Println("slice of a slice:", s[1:2])

	if slices.Equal(s, c) {
		fmt.Println("the two slices are equal:", s, c)
	}

	t := []string{"x", "y", "z"}
	fmt.Println("slices can also be direclty initlialized:", t)

	twoDslice := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D slices can have differen lengths:", twoDslice)

	//maps (Python dicts)
	m := make(map[string]int)
	m["m1"] = 1
	m["m2"] = 2

	fmt.Println("map m:", m, "with length:", len(m))

	delete(m, "m2")
	fmt.Println("map m:", m, "with length:", len(m))

	clear(m)
	fmt.Println("cleared map m:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map initialized in one line:", n)
	nCopy := make(map[string]int)
	maps.Copy(nCopy, n)
	if maps.Equal(n, nCopy) {
		fmt.Println("maps are copied with maps.Copy and compared with maps.Equals", n, nCopy)
	}

}
