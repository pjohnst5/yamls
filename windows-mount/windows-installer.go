package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)


func main() {
    fmt.Println("Hello, world.")
    fmt.Println("Printing C:\\test-pd")
    fmt.Println(filepath.Dir("C:\\test-pd"))

    fmt.Println("Writing file")
    message:= []byte("hello, gophers!")
    err := ioutil.WriteFile("C:\\test-pd\\hello.txt", message, 755)
    if err != nil {
        fmt.Println("Error: ", err)
    }
}