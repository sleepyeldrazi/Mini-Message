package main

import (
    "fmt"
    "net"
)

func main() {
    // Create listener
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close() // call later

    fmt.Println("Server listening on :8080")

    for {

        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        // Handle the connection in a goroutine
        go handleClient(conn)
        }
    }

func handleClient(conn net.Conn) {
    //close when done
    defer conn.Close()
    //create buffer to read data
    buffer := make([]byte, 1024)

   for {
        // read data
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Error raeding:", err)
            return
        }
    
        data := string(buffer[:n])
        fmt.Printf("%s\n", data)
    }
}
