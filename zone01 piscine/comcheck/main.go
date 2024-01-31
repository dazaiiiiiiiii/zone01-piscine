package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) >= 2 {
		messagee := "Alert!!!"
		for i := 0; i < len(arg); i++ {
			if arg[i] == "01" || arg[i] == "galaxy" || arg[i] == "galaxy 01" {
				fmt.Println(messagee)
				return
			}
		}
	}
}
