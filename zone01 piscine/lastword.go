package main

import (
    "fmt"
    "os"
)

func checkspace(r byte) bool {
    return (r == ' ' || r == '\t' || r == '\n')
}

func main() {
    arg := os.Args
    if len(arg) < 2 {
        return
    }
    start, end := 0, len(arg[1])-1
    for start <= end && checkspace(arg[1][start]) {
        start++
    }
    for start <= end && checkspace(arg[1][end]) {
        end--
    }
    if start > end {
        return
    }
    se := end
    for !checkspace(arg[1][se]) {
        se--
    }
    fmt.Println(arg[1][se+1 : end+1])
}