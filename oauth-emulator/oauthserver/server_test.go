package oauthserver

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKey(t *testing.T) {
	k, err := generateKey()
	assert.NoError(t, err)

	keyString, err := json.Marshal(k)

	server := New(string(keyString), "", nil)
	fmt.Println(server)
	fmt.Println(string(keyString))
}
