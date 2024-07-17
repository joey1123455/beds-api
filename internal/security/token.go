package security

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joey1123455/beds-api/internal/common"
	"github.com/joey1123455/beds-api/internal/config"
	db "github.com/joey1123455/beds-api/internal/db/sqlc"
)

const (
	tokenSize               = 256
	refreshTokenExpireTime  = 24
	generateTokenExpireTime = 1
)

var (
	validate             = "Validate"
	parseKey             = "parse key"
	ErrUnexpectedMethod  = errors.New("unexpected method")
	ErrDecoding          = errors.New("could not decode key")
	ErrCreatingParseKey  = errors.New("create: parse key")
	ErrCreatingSignToken = errors.New("create: sign token")
)

type Validator struct {
	Store  db.Store
	Config config.Config
}

func NewValidator(s db.Store, c config.Config) Validator {
	return Validator{
		Store:  s,
		Config: c,
	}
}

func (v Validator) GenerateJwtToken(userID string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(v.Config.JwtKey)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrDecoding, err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrCreatingParseKey, err)
	}

	now := time.Now().UTC()
	expires := now.Add(generateTokenExpireTime * time.Hour)

	claims := make(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = expires.Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["refresh"] = true

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrCreatingSignToken, err)
	}

	return token, nil
}

func (v Validator) GenerateRefreshToken(userID string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(v.Config.JwtKey)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrDecoding, err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrCreatingParseKey, err)
	}

	now := time.Now().UTC()
	expires := now.Add(refreshTokenExpireTime * time.Hour)

	claims := make(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = expires.Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["auth"] = true

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrCreatingSignToken, err)
	}

	return token, nil
}

// func (v Validator) TwoFactoreAuth(email string) {
// 	totp := gotp.NewDefaultTOTP(v.Config.OtpSecret)
// 	uri := totp.ProvisioningUri(email, "centralAuth")
// 	err := qrcode.WriteFile(uri, qrcode.Medium, tokenSize, "auth.png")
// 	if err != nil {
// 		logger.ErrorLogger(err)
// 	}
// 	qrterminal.GenerateWithConfig(uri, qrterminal.Config{
// 		Level:     qrterminal.L,
// 		Writer:    os.Stdout,
// 		BlackChar: qrterminal.BLACK,
// 		WhiteChar: qrterminal.WHITE,
// 	})
// }

func ValidateTokens(tokenStr, keyStr string) (jwt.MapClaims, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecoding, err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("%s:%s %w", validate, parseKey, err)
	}

	parsedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrUnexpectedMethod
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", validate, err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("%s: %w", validate, common.ErrTokenParsingFailed)
	}

	return claims, nil
}
