package main

import (
	"dip-golang/src/domain"
	"dip-golang/src/infrastructure"
	"fmt"
)

var userRepository domain.Repository

func main() {

	userRepository = infrastructure.NewTestRepository()
	user := getUserById(1)
	fmt.Printf("Find: %v  \n", user)
	createUser(1, "hoge")
}

func getUserById(id int) *domain.User {
	user := userRepository.FindById(id)
	return user
}

func createUser(id int, name string) {
	defer userRepository.Store(&domain.User{id, name})
}
