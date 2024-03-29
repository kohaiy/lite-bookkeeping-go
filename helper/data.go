package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/aidarkhanov/nanoid/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func Md5(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func EncodePassword(password string, slat string) string {
	return Md5(Md5(password) + slat)
}

func GenerateSlat() string {
	slat, _ := nanoid.New()
	return slat
}

type TokenPayload struct {
	ID   uint   `json:"id"`
	Slat string `json:"slat"`
}

type CustomClaims struct {
	jwt.StandardClaims
	TokenPayload
}

func GenerateToken(payload TokenPayload) string {
	tokenSecret := []byte(viper.GetString("tokenSecret"))
	tokenExpiredAt := viper.GetDuration("tokenExpiredAt")
	claims := CustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * tokenExpiredAt).Unix(),
		},
		payload,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(tokenSecret)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", ss)
}

func ParseToken(tokenString string) *CustomClaims {
	tokenSecret := []byte(viper.GetString("tokenSecret"))
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		return nil
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}

func GetNowTime() time.Time {
	return time.Now()
}
