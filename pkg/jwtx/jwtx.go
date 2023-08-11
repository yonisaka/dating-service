package jwtx

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yonisaka/dating-service/config"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    time.Time
	RtExpires    time.Time
}

type JWTImpl interface {
	CreateTokenJWT(userID int64) *TokenDetails
	RequestTokenJwt(token string) (interface{}, jwt.MapClaims, error)
}

type jwtx struct {
	cfg *config.Config
}

func NewJWT() JWTImpl {
	return &jwtx{}
}

func (j *jwtx) CreateTokenJWT(userID int64) *TokenDetails {
	// create ssid
	suuid := uuid.New().String()
	rsuuid := uuid.New().String()

	// create token
	tokenExpiry := time.Now().Add(time.Duration(j.cfg.JWT.JwtTokenExpiry) * time.Second)
	atClaims := jwt.MapClaims{
		"authorized":   true,
		"dtid":         suuid,
		"expiary_time": tokenExpiry,
		"user_id":      userID,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(j.cfg.JWT.JwtKey))

	refreshTokenExpiry := time.Now().Add(time.Duration(j.cfg.JWT.JwtRefreshTokenExpiry) * time.Second)
	ratClaims := jwt.MapClaims{
		"authorized":           true,
		"dtid":                 rsuuid,
		"refresh_expiary_time": refreshTokenExpiry,
		"user_id":              userID,
	}
	rat := jwt.NewWithClaims(jwt.SigningMethodHS256, ratClaims)
	refreshToken, _ := rat.SignedString([]byte(j.cfg.JWT.JwtKey))

	tokenDetails := &TokenDetails{
		AccessToken:  token,
		RefreshToken: refreshToken,
		AccessUuid:   suuid,
		RefreshUuid:  rsuuid,
		AtExpires:    tokenExpiry,
		RtExpires:    refreshTokenExpiry,
	}

	return tokenDetails
}

func (j *jwtx) RequestTokenJwt(authorizationHeader string) (interface{}, jwt.MapClaims, error) {
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, jwt.MapClaims{}, err
	}

	// get claim token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		suuid := fmt.Sprintf("%v", claims["dtid"])

		// validate token
		if !ok && !token.Valid {
			return nil, jwt.MapClaims{}, err
		}

		return suuid, claims, err
	}

	return nil, jwt.MapClaims{}, errors.New("token is not valid")
}
