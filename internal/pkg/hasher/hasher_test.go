package hasher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasher_GetHashFromStruct(t *testing.T) {
	password := "qwerty123"
	passHash, err := GetHashFromStruct(password)
	require.Nil(t, err)

	expectedHash := "823bac4e2352c6f5335766b154ede4ce4a38b0779004631b62aaae31eb5e1d45"
	require.Equal(t, expectedHash, passHash)
}
