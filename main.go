package main

import (
	"fmt"
	"jwt/jwt_token"
)

func main() {
	const userId = 12
	const nickname = "max"

	user := jwt_token.UserPayload { Nickname: nickname, UserId: userId }

	token, err := jwt_token.GenerateToken(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	//tokenFromRails := "eyJhbGciOiJIUzUxMiJ9.eyJkYXRhIjp7InVzZXJfaWQiOjEsIm5pY2tuYW1lIjoibmlja25hbWUifSwiZXhwIjoxNjMyMTg4Njk0LCJpYXQiOjE2MjQ0MTI2OTQsImlzcyI6IlRoZSBSaWNoLmlvIG9yIGh0dHBzOi8vd3d3LnRoZXJpY2guaW8vIn0.7iKhcQfC0oUnbnNYDDPTaZcX0g0iOeA_xYobQdzDaqBH46LeJVQlt_ivjr1V8w7VFO3aCp-PRdiidRAgDhJkvg"

	result, err := jwt_token.Verify(&token)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Decoded UserId is %d\n", result)
}
