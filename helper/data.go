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

var tokenSecret = []byte(viper.GetString("tokenSecret"))

func GenerateToken(payload TokenPayload) string {
	claims := CustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
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
