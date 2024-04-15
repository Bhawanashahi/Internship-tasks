
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// User struct represents the user data
type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

// UserDataStore represents the in-memory storage for user data
type UserDataStore struct {
	mu    sync.RWMutex
	users map[int]User
}

// NewUserDataStore initializes a new UserDataStore
func NewUserDataStore() *UserDataStore {
	return &UserDataStore{
		users: make(map[int]User),
	}
}

// AddUser adds a user to the UserDataStore
func (uds *UserDataStore) AddUser(user User) {
	uds.mu.Lock()
	defer uds.mu.Unlock()
	uds.users[user.UserID] = user
}

// GetUserByID retrieves a user by ID from the UserDataStore
func (uds *UserDataStore) GetUserByID(userID int) (User, bool) {
	uds.mu.RLock()
	defer uds.mu.RUnlock()
	user, ok := uds.users[userID]
	return user, ok
}

// GetAllUsers returns all users from the UserDataStore
func (uds *UserDataStore) GetAllUsers() []User {
	uds.mu.RLock()
	defer uds.mu.RUnlock()
	users := make([]User, 0, len(uds.users))
	for _, user := range uds.users {
		users = append(users, user)
	}
	return users
}

var userDataStore *UserDataStore

func init() {
	userDataStore = NewUserDataStore()
}

func main() {
	http.HandleFunc("/adduser", AddUserHandler)
	http.HandleFunc("/getuser", GetUserHandler)
	http.HandleFunc("/getallusers", GetAllUsersHandler)

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}

// AddUserHandler handles the addition of a new user
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userDataStore.AddUser(user)
	w.WriteHeader(http.StatusCreated)
}

// GetUserHandler handles retrieving a user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, ok := userDataStore.GetUserByID(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetAllUsersHandler handles retrieving all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	users := userDataStore.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}
