package main

import (
    "fmt"
    "os"
)

func main() {

     dirname := os.Args[1]

     d, err := os.Open(dirname)
     if err != nil {
         fmt.Println(err)
         os.Exit(1)
     }
     defer d.Close()

     files, err := d.Readdir(-1)
     if err != nil {
         fmt.Println(err)
         os.Exit(1)
     }

     fmt.Println("Reading "+ dirname)

     for _, file := range files {
         if file.Mode().IsRegular() {
             fmt.Println(file.Name(), file.Size(), "bytes")
         } else {
             fmt.Println("directory ", file.Name())
         }
     }
 }