package executil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseByEnv(t *testing.T) {
	tt := assert.New(t)

	os.Setenv("TEST", "1")
	tt.Equal("1", ParseByEnv("${TEST}"))
	tt.Equal("", ParseByEnv("${MISSING}"))
	tt.Equal("${IGNORE}", ParseByEnv("$${IGNORE}"))
}
