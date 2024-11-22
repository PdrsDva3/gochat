package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go-jwt-template/pkg/config"
	"time"
)

type JWTUtil struct {
	expireTimeOut time.Duration
	secret        string
}

func InitJWTUtil() JWTUtil {
	return JWTUtil{
		expireTimeOut: time.Duration(viper.GetInt(config.JWTExpire)) * time.Minute,
		secret:        viper.GetString(config.Secret),
	}
}

type userClaim struct {
	jwt.RegisteredClaims
	ID int
}

func (j JWTUtil) CreateToken(userID int) string {

	expiredAt := time.Now().Add(j.expireTimeOut)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expiredAt,
			},
		},
		ID: userID,
	})

	signedString, _ := token.SignedString([]byte(j.secret))

	return signedString
}

func (j JWTUtil) Authorize(tokenString string) (int, error) {
	var userClaim userClaim

	token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, nil
	}

	return userClaim.ID, nil
}
