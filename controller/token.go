package controller

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var ExpireTime = 3600 //token有效期一小时
var Secret = "chty_syq"//密钥

type JWTClaims struct {
	jwt.StandardClaims
	Username	string	`json: "username"`
	Password	string	`json: "password"`
}

func getToken(claims jwt.Claims)(string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil{
		return "", errors.New("Get token error!")
	}
	return signedToken, nil
}

func verifyToken(strToekn string) (error){
	token, err := jwt.ParseWithClaims(strToekn, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil{
		return errors.New("Parse token failed!")
	}
	_, ok := token.Claims.(*JWTClaims)
	if !ok{
		return errors.New("Please login!")
	}
	if err := token.Claims.Valid(); err != nil {
		return errors.New("Please login!")
	}
	return nil
}