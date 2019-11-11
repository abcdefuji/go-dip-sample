package infrastructure

import (
	"dip-golang/src/domain"
	"fmt"
)

type TestRepository struct{}

func (repo TestRepository) FindById(id int) *domain.User {
	return &domain.User{1, "test_user"}
}

func (repo TestRepository) Store(u *domain.User) {
	fmt.Printf("userId: %v \n", u.ID)
	fmt.Printf("userName: %v \n", u.Name)
	fmt.Println("stored success")
	return
}

func NewTestRepository() TestRepository {
	return TestRepository{}
}
