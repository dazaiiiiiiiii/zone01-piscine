package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) < 3 || len(arg) > 3 {
		return
	}
	res := ""
	for i := 0; i < len(arg[1]); i++ {
		for j := 0; j < len(arg[2]); j++ {
			if i < len(arg[1]) && arg[1][i] == arg[2][j] {
				res += string(arg[2][j])
				i++
			}
		}
	}
	if len(res) == len(arg[1]) {
		fmt.Println(res)
	}else{
		return
	}
}
