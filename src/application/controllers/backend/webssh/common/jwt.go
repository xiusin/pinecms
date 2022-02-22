package common

import (
	"github.com/golang-jwt/jwt"
	"time"
)


var jwtKet = []byte("config.Config.Jwt.Key")

type Claims struct {
	Userid int64
	jwt.StandardClaims
}

func ReleaseToken(id int64) (token string, err error) {
	expireTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &Claims{
		Userid: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "admin",
			Subject:   "user",
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenObj.SignedString(jwtKet)
	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKet, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
