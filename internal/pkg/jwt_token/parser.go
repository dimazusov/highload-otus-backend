package jwt_token

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const salt = "fawcoehfoiwefksdnfkbdksf"

type parser struct{}

var ErrWrongToken = errors.New("wrong token")

func NewParser() *parser {
	return &parser{}
}

func (m parser) Create(token string) (*jwtToken, error) {

}

func (m parser) Parse(token string) (*jwtToken, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.Wrapf(ErrWrongToken, "token parts must contains 3 part, given %d", len(parts))
	}

	decodedHeader, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, errors.Wrapf(err, "cannot decode header")
	}

	header := make(map[string]string)
	err = json.Unmarshal(decodedHeader, &header)
	if err != nil {
		return nil, err
	}

	decodedPayload, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, errors.Wrapf(err, "cannot decode payload")
	}

	payload := make(map[string]interface{})
	err = json.Unmarshal(decodedPayload, &payload)
	if err != nil {
		return nil, err
	}

	expectedSignature := m.getSignature(parts[0], parts[1])
	givenSignature, err := base64.StdEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}

	if expectedSignature != string(givenSignature) {
		return nil, errors.Wrapf(ErrWrongToken, "expected signature not equal given signature")
	}

	return &jwtToken{
		header:    header,
		payload:   payload,
		signature: expectedSignature,
	}, nil
}

func (m parser) getSignature(encodedHeader string, encodedPayload string) (hash string) {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s.%s%s", encodedHeader, encodedPayload, salt)))
	return fmt.Sprintf("%x", sum)[:32]
}
