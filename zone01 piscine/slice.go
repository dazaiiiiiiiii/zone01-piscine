package main

import (
	"fmt"

	"piscine"
)

func Slice(a []string, nbrs ...int) []string {
	start := nbrs[0]
	end := 0
	if len(nbrs) > 1 {
		end = nbrs[1]
	}
	if start < 0 {
		start += len(a)
	} else if end < 0 {
		end += len(a)
	}
	if start > end {
		return nil
	}
	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", piscine.Slice(a, 1))
	fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
	fmt.Printf("%#v\n", piscine.Slice(a, -3))
	fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
	fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
}
