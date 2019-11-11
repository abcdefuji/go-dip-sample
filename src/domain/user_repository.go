package domain

type Repository interface {
	Store(u *User)
	FindById(id int) *User
}
