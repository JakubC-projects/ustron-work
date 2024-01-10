package pages

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInde(t *testing.T) {
	wr := bytes.Buffer{}

	err := RenderLogin(&wr, Login{})

	assert.NoError(t, err)
	assert.NotEmpty(t, wr.String())

}
