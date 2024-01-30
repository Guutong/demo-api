package user

import "gorm.io/gorm"

type IUserRepository interface {
	NewUser(u *User) error
	GetUser() ([]User, error)
	DeleteUser(id int) error
	UpdateUser(id int, u *User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (us *UserRepository) NewUser(u *User) error {
	r := us.db.Create(u)
	return r.Error
}

func (us *UserRepository) GetUser() ([]User, error) {
	var users []User
	r := us.db.Find(&users)
	return users, r.Error
}

func (us *UserRepository) DeleteUser(id int) error {
	r := us.db.Delete(&User{}, id)
	return r.Error
}

func (us *UserRepository) UpdateUser(id int, u *User) error {
	r := us.db.Model(&User{}).Where("id = ?", id).Updates(u)
	return r.Error
}
