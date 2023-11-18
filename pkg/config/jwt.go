package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("aasfho[123r154456qf6qf]")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}