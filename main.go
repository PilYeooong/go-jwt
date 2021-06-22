package main

import (
	"fmt"
	"jwt/jwt_token"
)

func main() {
	const userId = "12"
	const nickname = "max"

	user := jwt_token.UserPayload { Nickname: nickname, UserId: userId }

	token, err := jwt_token.GenerateToken(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	result, err := jwt_token.Verify(&token)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Decoded UserId is %s\n", result)
}
