package jwt_token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

const TIMELIMIT = time.Minute * 60
var jwtSecret string

type VerifyResult struct {
	UserId   string
	Nickname string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file unexpected")
	}
	jwtSecret = os.Getenv("JWT_SECRET")
}

func GenerateToken(userId string, nickname string) (string, error) {
	payload := &jwt.MapClaims{
		"data": map[string]string{
			"nickname": nickname,
			"user_id":  userId,
		},
		"iss": "The Rich.io or https://www.therich.io/",
		"exp": time.Now().Add(TIMELIMIT).Unix(),
		"iat": time.Now().Unix(),
	}
	return Encode(payload)
}

func Encode(payload *jwt.MapClaims) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	token, err := at.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Decode(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetUserId(decodedData jwt.MapClaims) string {
	return decodedData["user_id"].(string)
}

func GetUserNickname(decodedData jwt.MapClaims) string {
	return decodedData["nickname"].(string)
}

func Verify(tokenString string) (string, error) {
	token, err := Decode(tokenString)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		decodedData := claims["data"].(map[string]interface{})
		userId := GetUserId(decodedData)

		return userId, nil
	}
	return "", err
}
