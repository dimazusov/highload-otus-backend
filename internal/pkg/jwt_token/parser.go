package jwt_token

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const algorithm = "SHA256"
const salt = "fawcoehfoiwefksdnfkbdksf"

type parser struct{}

var ErrWrongToken = errors.New("wrong token")

func NewParser() *parser {
	return &parser{}
}

func (m parser) Create(userID uint) (*jwtToken, error) {
	header := make(map[string]string)
	header["algorithm"] = algorithm

	payload := make(map[string]interface{})
	payload["userId"] = userID

	jsonHeader, err := json.Marshal(header)
	if err != nil {
		return nil, err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	encodedHeader := base64.StdEncoding.EncodeToString(jsonHeader)
	encodedPayload := base64.StdEncoding.EncodeToString(jsonPayload)

	signature := getSignature(encodedHeader, encodedPayload)

	return &jwtToken{
		header: header,
		payload: payload,
		signature: signature,
	}, nil
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

	expectedSignature := getSignature(parts[0], parts[1])
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

func getSignature(encodedHeader string, encodedPayload string) (hash string) {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s.%s%s", encodedHeader, encodedPayload, salt)))
	return fmt.Sprintf("%x", sum)[:32]
}
