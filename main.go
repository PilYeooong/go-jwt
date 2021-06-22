package main

import (
	"fmt"
	"jwt/jwt_token"
)

func main() {
	userId := "1"
	nickname := "max"

	token, err := jwt_token.GenerateToken(userId, nickname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	result, err := jwt_token.Verify(token)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Decoded UserId is %s\n", result)
}
