package main

// import (
//     "fmt"
//     "net"
// )

// func main() {
//     addr := net.UDPAddr{
//         Port: 8080,
//         IP:   net.ParseIP("127.0.0.1"),
//     }

//     conn, err := net.ListenUDP("udp", &addr)
//     if err != nil {
//         fmt.Println("Error listening:", err.Error())
//         return
//     }
//     defer conn.Close()

//     fmt.Println("Server is listening on port 8080...")

//     for {
//         buffer := make([]byte, 1024)
//         n, clientAddr, err := conn.ReadFromUDP(buffer)
//         if err != nil {
//             fmt.Println("Error reading:", err.Error())
//             continue
//         }
//         message := string(buffer[:n])
//         fmt.Printf("Message received from client: %s\n", message)

//         newMessage := "Hello, " + message
//         _, err = conn.WriteToUDP([]byte(newMessage), clientAddr)
//         if err != nil {
//             fmt.Println("Error writing:", err.Error())
//             return
//         }
//     }
// }
