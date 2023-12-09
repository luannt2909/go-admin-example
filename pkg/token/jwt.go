package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Tokenizer interface {
	Generate(claim Claim) (string, error)
	Parse(token string) (Claim, error)
}
type jwtTokenizer struct {
	signingKey []byte
}

func (j *jwtTokenizer) Generate(claim Claim) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"claim": claim,
		"iat":   time.Now(),
	})
	tokenString, err = token.SignedString(j.signingKey)
	return
}

func (j *jwtTokenizer) Parse(tokenStr string) (claim Claim, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.signingKey, nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("claim is invalid")
		return
	}
	c, ok := claims["claim"]
	if !ok {
		err = errors.New("claim is invalid")
		return
	}
	cBytes, err := json.Marshal(c)
	if err != nil {
		return
	}
	err = json.Unmarshal(cBytes, &claim)
	return
}

func NewJwtTokenizer(signingKey []byte) Tokenizer {
	return &jwtTokenizer{signingKey: signingKey}
}
