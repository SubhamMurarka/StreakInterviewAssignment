package UserStore

import (
	"fmt"
	"log"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// Representing a user
type User struct {
	UserName       string
	HashedPassword string
	AccessToken    string
}

// Inmemory DB to store the user
type UserStore struct {
	store map[string]*User
	mu    sync.RWMutex
}

// Interface containing all the functions for Inmemory DB
type UserStoreInterface interface {
	Save(username, password string) (*User, error)
	Find(username string) *User
	IsCorrectPassword(username, password string) bool
}

// Creating New object for the UserStore
func NewUserStore() UserStoreInterface {
	return &UserStore{
		store: make(map[string]*User),
	}
}

// Saving User to the database
func (u *UserStore) Save(username, password string) (*User, error) {
	//Hashing user password
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error Hashing Password")
		return nil, err
	}

	user := &User{
		UserName:       username,
		HashedPassword: string(hashedpassword),
	}

	//Acquiring Lock over the inmemory map
	u.mu.Lock()
	defer u.mu.Unlock()

	//checking if user already exist
	if u.store[username] != nil {
		return nil, fmt.Errorf("user already exists")
	}

	u.store[username] = user

	for _, val := range u.store {
		fmt.Println(val.UserName)
		// fmt.Printf("Username %s : ", val.UserName)
	}

	return user, nil
}

// fetching User from Inmemory DB
func (u *UserStore) Find(username string) *User {
	//Acquiring Read Lock over the inmemory map
	u.mu.RLock()
	defer u.mu.RUnlock()

	//checking if user exist or not
	if u.store[username] == nil {
		return nil
	}

	return u.store[username]
}

func (u *UserStore) IsCorrectPassword(username, password string) bool {
	//Acquiring Read Lock over the inmemory map
	u.mu.RLock()
	defer u.mu.RUnlock()

	if u.store[username] == nil {
		log.Printf("User Doesnot Exist")
		return false
	}

	user := u.store[username]

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		log.Printf("Wrong Password")
		return false
	}

	return true
}
