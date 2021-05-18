package main

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/pascaldekloe/jwt"
)

var JWTPrivateKey ed25519.PrivateKey
var JWTPublicKey ed25519.PublicKey

func initJWT() {
	seed := Config.Seed
	if seed == "" {
		return
	}
	for len(seed) < ed25519.SeedSize {
		seed = seed + seed
	}
	if len(seed) > ed25519.SeedSize {
		seed = seed[:ed25519.SeedSize]
	}

	JWTPrivateKey = ed25519.NewKeyFromSeed([]byte(seed))
	JWTPublicKey = []byte(JWTPrivateKey)[32:]
}

func createUserToken(groups []string) ([]byte, error) {
	var claims jwt.Claims
	claims.Subject = "user"
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Set = make(map[string]interface{}, len(groups))
	for _, x := range groups {
		claims.Set[x] = true
	}
	return claims.EdDSASign(JWTPrivateKey)
}

func verifyUserToken(token []byte) (map[string]bool, error) {
	claims, err := jwt.EdDSACheck(token, JWTPublicKey)
	if err != nil {
		return nil, err
	}
	if !claims.Valid(time.Now()) {
		return nil, fmt.Errorf("credential time constraints exceeded")
	}
	if claims.Subject != "user" {
		return nil, fmt.Errorf("wrong claims subject")
	}

	out := make(map[string]bool)
	for k, v := range claims.Set {
		out[k] = v.(bool)
	}
	return out, nil
}
