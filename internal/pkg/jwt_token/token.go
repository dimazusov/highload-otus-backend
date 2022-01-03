package jwt_token

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type jwtToken struct {
	header    map[string]string
	payload   map[string]interface{}
	signature string
}

func (m *jwtToken) SetUserID(userID uint) {
	m.payload["userId"] = userID
}

func (m *jwtToken) GetUserID() uint {
	return uint(m.payload["userId"].(float64))
}

func (m *jwtToken) ToString() (string, error) {
	jsonHeader, err := json.Marshal(m.header)
	if err != nil {
		return "", err
	}

	jsonPayload, err := json.Marshal(m.payload)
	if err != nil {
		return "", err
	}

	encodedHeader := base64.StdEncoding.EncodeToString(jsonHeader)
	encodedPayload := base64.StdEncoding.EncodeToString(jsonPayload)
	encodedSignature := base64.StdEncoding.EncodeToString([]byte(getSignature(encodedHeader, encodedPayload)))

	return fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, encodedSignature), nil
}
