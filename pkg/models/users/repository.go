package users

import (
	"Api/pkg/service"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) GetHashedPassword(email string, password string) (int, error) {
	var user Users

	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return 0, err
	}
	err = service.BcryptHasher{}.CheckPasswordHash(password, user.Password)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *UserRepository) CreateUser(email string, password string, name string) bool {
	var user Users
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		if user.Password != "" {
			return false
		}
		user.Password = password
		user.Name = name
		u.DB.Save(&user)
		return true
	}
	user.Email = email
	user.Password = password
	user.Name = name
	u.DB.Select("email", "password", "name").Create(&user)
	return true
}

func (u *UserRepository) GetEmail(id int) (string, error) {
	var user Users

	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (u *UserRepository) GetUser(email, password string) (user Users) {
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		user.Email = email
		user.Password = password
		user.Name = "google"
		u.DB.Select("email", "password", "name").Create(&user)
	} else {
		return
	}
	return user
}
