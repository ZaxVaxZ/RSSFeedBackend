package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authorization header found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authorization header (First part should be \"ApiKey\")")
	}
	if len(vals[1]) != 64 {
		return "", fmt.Errorf("malformed authorization header (Invalid API Key) %v", vals[1])
	}

	return vals[1], nil
}
