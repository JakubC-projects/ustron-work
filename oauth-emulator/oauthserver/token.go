package oauthserver

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

func (s *Server) newToken(payload any, duration time.Duration) (string, error) {
	builder := jwt.NewBuilder().
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(duration))

	builder, err := addStructToTokenClaims(builder, payload)
	if err != nil {
		return "", fmt.Errorf("cannot create token: %w", err)
	}

	tok, err := builder.Build()
	if err != nil {
		return "", fmt.Errorf("cannot create token: %w", err)
	}

	return s.signToken(tok)
}

func (s *Server) signToken(tok jwt.Token) (string, error) {
	tokenStr, err := jwt.Sign(tok,
		jwa.SignatureAlgorithm(s.signingKey.Algorithm()),
		s.signingKey,
	)

	if err != nil {
		return "", fmt.Errorf("cannot sign token: %w", err)
	}

	return string(tokenStr), nil
}

func addStructToTokenClaims(b *jwt.Builder, data any) (*jwt.Builder, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("canont convert token claims to json string: %w", err)
	}

	var claims map[string]any

	err = json.Unmarshal(jsonBytes, &claims)
	if err != nil {
		return nil, fmt.Errorf("canont unmarshal claims into struct: %w", err)
	}

	for claim, val := range claims {
		b = b.Claim(claim, val)
	}
	return b, nil
}
