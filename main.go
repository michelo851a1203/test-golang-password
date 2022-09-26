package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateCompare(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func EncryptPassword(inputPassword string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func main() {
	message := "testing hello"
	result, err := EncryptPassword(message)
	if err != nil {
		log.Fatalf("result error")
		return
	}
	fmt.Println(result)
	isComparedOK := GenerateCompare(result, message)
	fmt.Println(isComparedOK)
}
