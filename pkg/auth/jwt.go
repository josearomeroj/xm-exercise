package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type userClaims struct {
	jwt.StandardClaims
	Id uuid.UUID `json:"id"`
}

type JWTManager struct {
	privateKey    *rsa.PrivateKey
	publicKey     any
	tokenDuration time.Duration
}

func NewJWTManager(privateKey string, publicKey string, tokenDuration time.Duration) (*JWTManager, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("error loading private key")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	block, _ = pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("error loading public key")
	}
	pk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &JWTManager{key, pk, tokenDuration * time.Millisecond}, nil
}

func (m *JWTManager) Generate(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, &userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.tokenDuration).Unix(),
		},
		Id: id,
	})
	return token.SignedString(m.privateKey)
}

func (m *JWTManager) Verify(t string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(t, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return m.publicKey, nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid token claims")
	}

	return claims.Id, nil
}
