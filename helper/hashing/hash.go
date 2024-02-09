package hashing

import "golang.org/x/crypto/bcrypt"

type HashInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}

type Hash struct {
}

func NewHash() HashInterface {
	return &Hash{}
}

func (h *Hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *Hash) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
