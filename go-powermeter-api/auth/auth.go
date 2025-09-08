package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJwtKey() []byte {
	// Ambil nilai dari environment variable yang bernama JWT_SECRET
	key := os.Getenv("JWT_SECRET")
	// if key == "" {
	//     // Jika tidak ada, gunakan nilai default (hanya untuk development)
	// 	return []byte("kunci_default_jika_env_tidak_ditemukan")
	// }
	return []byte(key)
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Gunakan fungsi untuk mendapatkan kunci
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJwtKey())
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Gunakan fungsi untuk mendapatkan kunci
		return getJwtKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}
