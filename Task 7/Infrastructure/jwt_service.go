package Infrastructure

import (
	"fmt"
	"os"
	"task/Domain"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(user *Domain.User, secret string) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * 24 * 30).Unix()
	claims := &Domain.JwrCustonClaims{
		UserName: user.UserName,
		Role:     user.Role,
		
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err = token.SignedString([]byte(secret))
	if err != nil {
	
		return "", err
	}
	return accessToken, nil
}

func VerifyAccessToken(accessToken string, secret string) (bool, error) {
	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(accessToken string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("couldn't parse claims")

	}
	
	return claims, nil
}

