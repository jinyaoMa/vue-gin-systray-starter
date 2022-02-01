package claims

import (
	"app/config"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

type Claims struct {
	jwt.RegisteredClaims
	UserInfo UserInfo
}

type UserInfo struct {
	ID         uint
	Name       string
	Account    string
	Permission string
}

func New(config *config.Claims, host string, userInfo UserInfo) *Claims {
	now := time.Now()
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  config.Issuer,
			Subject: host,
			Audience: []string{
				config.Issuer,
				userInfo.Account,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: now.AddDate(0, 0, config.DaysToExpire),
			},
			NotBefore: &jwt.NumericDate{
				Time: now,
			},
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
			ID: strconv.FormatUint(uint64(userInfo.ID), 10),
		},
		UserInfo: userInfo,
	}
	return claims
}

func Parse(token string, key string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

func (c *Claims) ToToken(key string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	return t.SignedString([]byte(key))
}
