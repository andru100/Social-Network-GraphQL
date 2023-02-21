package social

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)

func Chkauth(tknStr *string) (*model.Authd, error)  { // checks for authentication

	var jwtKey = []byte("AllYourBase")

	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(*tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	var auth model.Authd

	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}

	auth.AuthdUser = claims.Username

	return &auth, err
}