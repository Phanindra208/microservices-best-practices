package utils

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// JWTPayload contains JWT payload.
type JWTPayload map[string]interface{}

// GetJWTPayload parses given JWT and returns payload object.
func GetJWTPayload(jwt string) (JWTPayload, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) < 3 {
		return nil, fmt.Errorf("jwt is missing parts")
	}

	payloadStr := parts[1]
	if l := len(payloadStr) % 4; l > 0 {
		payloadStr += strings.Repeat("=", 4-l)
	}

	b, err := base64.URLEncoding.DecodeString(payloadStr)
	if err != nil {
		return nil, err
	}

	var payload JWTPayload
	if err := JSON.Unmarshal(b, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}

// GetMemberIDFromToken finds member ID from token.
func GetMemberIDFromToken(token string) (string, bool) {
	payload, err := GetJWTPayload(token)
	if err != nil {
		return "", false
	}
	for key, value := range payload {
		if strings.Contains(key, "gravty") && strings.Contains(key, "memberId") {
			return value.(string), true
		}
	}
	return "", false
}

// GetEmailFromToken finds user ID from token.
func GetEmailFromToken(token string) (string, bool) {
	payload, err := GetJWTPayload(token)
	if err != nil {
		return "", false
	}
	for key, value := range payload {
		if strings.EqualFold(key, "") {
			return value.(string), true
		}
	}
	return "", false
}
