package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    var user string = "Koko"
    // connect to server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }

    defer conn.Close()
   
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Println("Enter message: ")
        // data to send
   
        if scanner.Scan() {
            data := scanner.Text()
            if err:= scanner.Err(); err!= nil {
                fmt.Println("Error reading: ", err)
                return
            }

        // Send data
        _, err = conn.Write([]byte(user + ": " + data))
        if err != nil{
            fmt.Println("Error sending:", err)
            return
        }
    }
}
}
