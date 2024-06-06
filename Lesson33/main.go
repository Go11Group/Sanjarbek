package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/hello", hello)
// 	mux.HandleFunc("/name", name)
// 	mux.HandleFunc("/surname", surname)

// 	// Listen on all interfaces on port 8080
// 	err := http.ListenAndServe(":8080", mux)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.Write([]byte("This method is not GET"))
// 		return
// 	}
// 	fmt.Println("URL:", r.URL)
// 	fmt.Println("Host:", r.Host)
// 	fmt.Println("Method:", r.Method)

// 	n, err := w.Write([]byte("IN HELLOOOO"))
// 	if err != nil {
// 		fmt.Println(err, n)
// 	}
// }

// func name(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.Write([]byte("This method is not GET"))
// 		return
// 	}
// 	n, err := w.Write([]byte("My name is Sanjarbek"))
// 	if err != nil {
// 		fmt.Println(err, n)
// 	}
// }

// func surname(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.Write([]byte("This method is not GET"))
// 		return
// 	}
// 	n, err := w.Write([]byte("My surname is Abduraxmonov"))
// 	if err != nil {
// 		fmt.Println(err, n)
// 	}
// }
