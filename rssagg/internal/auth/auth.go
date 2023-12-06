package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	auth_header := headers.Get("Authorization")
	if auth_header == "" {
		return "", errors.New("no authentication header found")
	}

	splits := strings.Split(auth_header, " ")
	if len(splits) != 2 {
		return "", errors.New("malformed auth header")
	}

	if splits[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return splits[1], nil
}
