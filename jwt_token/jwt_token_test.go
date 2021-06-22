package jwt_token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const USERID = "1"
const NICKNAME = "MAX"

func TestEncode(t *testing.T) {
	user := UserPayload{UserId: USERID, Nickname: NICKNAME}
	payload := &jwt.MapClaims{
		"data": map[string]string{
			"nickname": user.Nickname,
			"user_id":  user.UserId,
		},
		"iss": "The Rich.io or https://www.therich.io/",
		"exp": time.Now().Add(TIMELIMIT).Unix(),
		"iat": time.Now().Unix(),
	}
	result, err := Encode(payload)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGenerateToken(t *testing.T) {
	user := UserPayload{UserId: USERID, Nickname: NICKNAME}
	token, err := GenerateToken(&user)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotNil(t, token)
}

func TestDecode(t *testing.T) {
	user := UserPayload{UserId: USERID, Nickname: NICKNAME}
	token, err := GenerateToken(&user)
	if err != nil {
		fmt.Println(err)
	}
	decodedToken, err := Decode(&token)
	if err != nil {
		fmt.Println(err)
	}
	assert.NotNil(t, decodedToken)
}

func TestVerify(t *testing.T) {
	user := UserPayload{UserId: USERID, Nickname: NICKNAME}
	token, err := GenerateToken(&user)
	if err != nil {
		fmt.Println(err)
	}
	verifyResult, err := Verify(&token)
	assert.NotNil(t, verifyResult)
	assert.Equal(t, verifyResult, USERID)
}

func TestGetUserId(t *testing.T) {
	//user := UserPayload{UserId: USERID, Nickname: NICKNAME}
	//token, err := GenerateToken(&user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//verifyResult, err := Verify(&token)
}