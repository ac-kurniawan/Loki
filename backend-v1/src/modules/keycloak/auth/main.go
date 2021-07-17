package keycloak

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	"github.com/dgrijalva/jwt-go/v4"
)

type IKeycloak interface {
	newToken() (*gocloak.JWT, error)
	Login(username string, password string) (*gocloak.JWT, error)
	ValidateToken(token string) (*gocloak.RetrospecTokenResult, error)
	DecodeToken(token string) (*jwt.Token, error)
}

type KeycloakTokenClaims struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type Module struct {
	client gocloak.GoCloak
	clientId string
	clientSecret string
	realmName string
	realmHost string
}

func (m Module) DecodeToken(token string) (*jwt.Token, error) {
	return m.client.DecodeAccessTokenCustomClaims(context.Background(), token, m.realmName,
		"account", &KeycloakTokenClaims{})
}

func (m Module) newToken() (*gocloak.JWT, error) {
	token, err := m.client.LoginClient(context.Background(), m.clientId,
		m.clientSecret, m.realmName)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (m Module) Login(username string, password string) (*gocloak.JWT, error){
	return m.client.Login(context.Background(), m.clientId,
		m.clientSecret, m.realmName, username, password)
}

func (m Module) ValidateToken(token string) (*gocloak.RetrospecTokenResult, error) {
	return m.client.RetrospectToken(context.Background(), token, m.clientId,
		m.clientSecret, m.realmName)
}

func NewKeycloak(realmHost string, realmName string,
	clientId string, clientSecret string) IKeycloak {
	client := gocloak.NewClient(realmHost)
	return &Module{
		client: client,
		clientId: clientId,
		clientSecret: clientSecret,
		realmName: realmName,
		realmHost: realmHost,
	}
}