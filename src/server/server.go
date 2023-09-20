package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
)

func main() {
    // log file
    logFileName := "server.log"

    file, err := os.Open(logFileName)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("Log doesn't exists: ", err)
        } else {
            fmt.Println("Error opening file: ", err)
        }
    }

    // scanner to read lines
    scanner := bufio.NewScanner(file)

    // read each line
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil{
        fmt.Println("Error reading: ", err)
    }

    file.Close() // explicitly close since we will write later on

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
        go handleClient(conn, logFileName)
        }
    }

func handleClient(conn net.Conn, fileName string) {
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
        go writeToLog(fileName, data)
    }
}

func writeToLog(fileName, data string)  {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        fmt.Println("Error: ", err)
        return 
    }
   
    defer file.Close()

    _, writeErr := file.WriteString(data + "\n")
    if writeErr != nil{
        fmt.Println("Error writing to log: ", err)
        return
    }

    return


}
