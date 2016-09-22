package main

import (
    "net"
    "os"
    "fmt"
)

func main() {
    service := ":1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        handleClient(conn)
        conn.Close()
    }
}

func handleClient(conn net.Conn) {
    var buff [512]byte
    for {
        n, err := conn.Read(buff[0:]) // n is the number of bytes recieved
        if err != nil { // error recieveing
            return
        }
        fmt.Println(string(buff[0:]))
        _, err2 := conn.Write(buff[0:n])
        if err2 != nil { // unable to send
            return
        }
    }
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
