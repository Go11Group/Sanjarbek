package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    serverAddr := net.UDPAddr{
        Port: 8080,
        IP:   net.ParseIP("127.0.0.1"),
    }

    conn, err := net.DialUDP("udp", nil, &serverAddr)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    go listenForMessages(conn)

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter message: ")
        message, _ := reader.ReadString('\n')

        _, err = conn.Write([]byte(message))
        if err != nil {
            fmt.Println("Error writing to UDP connection:", err)
            return
        }
    }
}

func listenForMessages(conn *net.UDPConn) {
    buffer := make([]byte, 1024)
    for {
        n, _, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error reading from UDP connection:", err)
            return
        }
        fmt.Println("Broadcast message:", string(buffer[:n]))
    }
}
