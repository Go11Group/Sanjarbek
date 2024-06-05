package main

import (
    "fmt"
    "net"
    "strings"
    "sync"
)

var clients = make(map[string]*net.UDPAddr)
var mu sync.Mutex

func main() {
    addr := net.UDPAddr{
        Port: 8080,
        IP:   net.ParseIP("127.0.0.1"),
    }

    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Println("Error setting up listener:", err)
        return
    }
    defer conn.Close()
    fmt.Println("Server is listening on port 8080...")

    buffer := make([]byte, 1024)
    for {
        n, clientAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error reading from UDP connection:", err)
            continue
        }

        message := string(buffer[:n])
        fmt.Printf("Received message from %s: %s\n", clientAddr.String(), message)

        mu.Lock()
        clients[clientAddr.String()] = clientAddr
        mu.Unlock()

        broadcastMessage(conn, clientAddr, message)
    }
}

func broadcastMessage(conn *net.UDPConn, senderAddr *net.UDPAddr, message string) {
    mu.Lock()
    defer mu.Unlock()

    for addrStr, addr := range clients {
        if addrStr != senderAddr.String() {
            response := fmt.Sprintf("Message from %s: %s", senderAddr.String(), strings.TrimSpace(message))
            _, err := conn.WriteToUDP([]byte(response), addr)
            if err != nil {
                fmt.Println("Error writing to UDP connection:", err)
            }
        }
    }
}
