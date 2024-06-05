package main

// import (
//     "bufio"
//     "fmt"
//     "net"
//     "os"
// )

// func main() {
//     // Create a UDP address to connect to the server
//     serverAddr := net.UDPAddr{
//         Port: 8080,
//         IP:   net.ParseIP("127.0.0.1"),
//     }

//     // Create a UDP connection
//     conn, err := net.DialUDP("udp", nil, &serverAddr)
//     if err != nil {
//         fmt.Println("Error connecting to server:", err)
//         return
//     }
//     defer conn.Close()

//     reader := bufio.NewReader(os.Stdin)
//     for {
//         // Read message from the user
//         fmt.Print("Enter message: ")
//         message, _ := reader.ReadString('\n')

//         // Send message to the server
//         _, err = conn.Write([]byte(message))
//         if err != nil {
//             fmt.Println("Error writing to UDP connection:", err)
//             return
//         }

//         // Read response from the server
//         buffer := make([]byte, 1024)
//         n, _, err := conn.ReadFromUDP(buffer)
//         if err != nil {
//             fmt.Println("Error reading from UDP connection:", err)
//             return
//         }

//         fmt.Println("Response from server:", string(buffer[:n]))
//     }
// }
