package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/create", createHandler)
	http.HandleFunc("/api/update/", updateHandler)
	http.HandleFunc("/api/delete/", deleteHandler)
	http.HandleFunc("/api/status", statusHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/user/", userHandler)
	http.HandleFunc("/api/logout", logoutHandler)
	http.HandleFunc("/api/products", productsHandler)
	http.HandleFunc("/api/profile", profileHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	response := fmt.Sprintf("Hello, %s", name)
	_, _ = w.Write([]byte(response))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	response := fmt.Sprintf("Created resource with data: %s", string(body))
	_, _ = w.Write([]byte(response))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/update/")
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	response := fmt.Sprintf("Updated resource %s with data: %s", id, string(body))
	_, _ = w.Write([]byte(response))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/delete/")
	response := fmt.Sprintf("Deleted resource %s", id)
	_, _ = w.Write([]byte(response))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := "Server status: OK"
	_, _ = w.Write([]byte(response))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	var credentials map[string]string
	_ = json.Unmarshal(body, &credentials)
	response := fmt.Sprintf("Login successful for user: %s", credentials["username"])
	_, _ = w.Write([]byte(response))
}


func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/user/")
	response := fmt.Sprintf("User details for ID %s", id)
	_, _ = w.Write([]byte(response))
}


func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := "Logout successful"
	_, _ = w.Write([]byte(response))
}


func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := "List of products"
	_, _ = w.Write([]byte(response))
}


func profileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	response := fmt.Sprintf("Profile updated with data: %s", string(body))
	_, _ = w.Write([]byte(response))
}
