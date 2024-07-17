package common

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(userID, jwtKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()
	expires := now.Add(2 * time.Hour)

	claims := make(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = expires.Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["refresh"] = true

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func GenerateRefreshToken(userID, jwtKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()
	expires := now.Add(24 * time.Hour)

	claims := make(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = expires.Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["auth"] = true

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ValidateTokens(tokenStr, keyStr string) (jwt.MapClaims, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims, nil
}
