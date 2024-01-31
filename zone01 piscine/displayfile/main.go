package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	// arg := os.Args
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Print(string(content))
}
