package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {
}

func (hash *Hash) Make(text string) (string, error) {

	text_byte := []byte(text)

	hashed, err := bcrypt.GenerateFromPassword(text_byte, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), err
}

func (hash *Hash) Check(hashed string, text string) bool {

	hashed_byte := []byte(hashed)
	text_byte := []byte(text)

	err := bcrypt.CompareHashAndPassword(hashed_byte, text_byte)

	if err == nil {
		return true
	}

	return false

}
