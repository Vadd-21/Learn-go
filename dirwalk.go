package main

import (
    "fmt"
    "path/filepath"
    "os"
)

func print_file(path string, info os.FileInfo, err error) error {
    if err != nil {
        return nil
    }
    fmt.Println(path) // prints every file it comes across
    return nil
}

func main() {
    dir := os.Args[1] // walks every directory from the provided up
    filepath.Walk(dir, print_file) 
    
}