package oauthserver

import "github.com/google/uuid"

type TokenPayload struct {
	Issuer    string    `json:"iss"`
	Audience  string    `json:"aud"`
	Subject   string    `json:"sub"`
	PersonUid uuid.UUID `json:"https://login.bcc.no/claims/personUid"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type OidcDiscovery struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	JwksUri               string `json:"jwks_uri"`
}
