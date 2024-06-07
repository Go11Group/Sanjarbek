package handler

import (
	"encoding/json"
	"module/model"
	"net/http"
	"strconv"
	"strings"
)

func (u *Users) UserCreate(w http.ResponseWriter, r *http.Request) {
	user := model.Users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = u.User.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func (u *Users) UserRead(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/ReadUsers/"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    users, err := u.User.GetUsers()
    if err != nil {
        http.Error(w, "Failed to get users", http.StatusInternalServerError)
        return
    }

    var user model.Users
    for _, u := range users {
        if u.Id == id {
            user = u
            break
        }
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(user)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}



func (u *Users) UserUpdate(w http.ResponseWriter, r *http.Request) {
	user := model.Users{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = u.User.UpdateUser(user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User updated successfully"))
}

func (u *Users) UserDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/DeleteUsers/"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = u.User.DeleteUser(id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User deleted successfully"))
}
