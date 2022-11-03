package service

import "golang.org/x/crypto/bcrypt"

type HasherI interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) error
}

type BcryptHasher struct {
	Cost int
}

func (h BcryptHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.Cost)
	return string(bytes), err
}

func (h BcryptHasher) CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
