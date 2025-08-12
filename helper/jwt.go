package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Claims struct {
	Email string `json:"email"`
	Nama  string `json:"nama"`
	jwt.StandardClaims
}

func GenerateJWT(email, nama string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	jwtKey := os.Getenv("JWT_SECRET")
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: email,
		Nama:  nama,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtKey))
}
