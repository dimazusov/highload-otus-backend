package jwt_token

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_Create(t *testing.T) {
	userID := uint(2312)
	token, err := NewParser().Create(userID)
	require.Nil(t, err)
	require.Equal(t, userID, token.GetUserID())

	expectedToken := "eyJhbGdvcml0aG0iOiJTSEEyNTYifQ==.eyJ1c2VySWQiOjIzMTJ9.OGE0NDA3NDg0MGQzNGU4Mzk5YzRhZGE0OWU2OGZlMTk="
	givenToken, err := token.ToString()
	require.Nil(t, err)
	require.Equal(t, expectedToken, givenToken)
}

func TestParser_Parse(t *testing.T) {
	expectedToken := "eyJhbGdvcml0aG0iOiJTSEEyNTYifQ==.eyJ1c2VySWQiOjIzMTJ9.OGE0NDA3NDg0MGQzNGU4Mzk5YzRhZGE0OWU2OGZlMTk="

	p := NewParser()
	token, err := p.Parse(expectedToken)
	require.Nil(t, err)
	log.Println(token.GetUserID())
}

func TestParser_getSignature(t *testing.T) {
	header := `{"typ": "JWT","alg": "HS256"}`
	payload := `{"id": "1337","username": "bizone","iat": 1594209600,"role":"user"}`

	givenSignature := getSignature(header, payload)
	expectedSignature := "c9a3593c5767444815bc460e117a90ae"

	require.Equal(t, expectedSignature, givenSignature)
}
