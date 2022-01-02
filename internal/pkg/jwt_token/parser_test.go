package jwt_token

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	token := "eyJ0eXBlIjoiSldUIiwiYWxnIjoiSFMyNTYifQ==.eyJpZCI6IjEzMzciLCJ1c2VybmFtZSI6ImJpem9uZSIsImlhdCI6MTU5NDIwOTYwMCwicm9sZSI6InVzZXIifQ==.YmUwZTFlNTRiZmRiY2FjYmFmOWM2NWE1ZjQzZGI5MjA"

	p := NewParser()
	jwtToken, err := p.Parse(token)
	// require.Nil(t, err)
	log.Println(err)

	log.Println(jwtToken)
}

func TestParser_getSignature(t *testing.T) {
	header := `{"typ": "JWT","alg": "HS256"}`
	payload := `{"id": "1337","username": "bizone","iat": 1594209600,"role":"user"}`

	givenSignature := NewParser().getSignature(header, payload)
	expectedSignature := "c9a3593c5767444815bc460e117a90ae"

	require.Equal(t, expectedSignature, givenSignature)
}