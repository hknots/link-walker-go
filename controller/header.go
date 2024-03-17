package controller

import (
	"errors"
	"net/http"
	"strings"
)

func GetAuthorizationToken(request *http.Request) (string, error) {
    authHeader := request.Header.Get("Authorization")
    if authHeader == "" {
        return "", errors.New("authorization header is missing")
    }
    if !strings.HasPrefix(authHeader, "Bearer ") {
        return "", errors.New("authorization header must start with Bearer")
    }

    token := strings.TrimPrefix(authHeader, "Bearer ")
    return token, nil
}

type AuthHeader struct {
	Authorization string
}
