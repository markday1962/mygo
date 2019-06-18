package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p1 := "qwerty"
	DefaultCost := 10

	bs, err := bcrypt.GenerateFromPassword([]byte(p1), DefaultCost)
	if err != nil {
		fmt.Println("Error caught generating password hash", err)
	}
	fmt.Println("Password hash:", bs)
	fmt.Printf("Type of password hash is %T\n", bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte(p1))
	if err != nil {
		fmt.Println("Incorrect password", err)
	} else {
		fmt.Println(p1, "Correct password")
	}

	p2 := "asdfghj"

	err = bcrypt.CompareHashAndPassword(bs, []byte(p2))
	if err != nil {
		fmt.Println(p2, "Incorrect password:", err)
	} else {
		fmt.Println("Correct password")
	}
}
