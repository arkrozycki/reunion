package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {

	str := "a test"
	hash := Hash(str)
	shouldBe := "9939b05dd1a3763f5f856e065d277190d648994f"

	assert.Equal(t, hash, shouldBe, "The hash should equal")
}
