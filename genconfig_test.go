package loadenv

import (
	"crypto/rand"
	"encoding/base32"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredExists(t *testing.T) {
	type testConfig struct {
		FOO string `env:"FOO" required:"true" default:"default value"`
	}

	// Save initial state and setup test env
	initialState := os.Getenv("FOO")
	expected := genToken(16)
	os.Setenv("FOO", expected)

	defer func() {
		r := recover()
		assert.Nil(t, r)
	}()

	actual := GenConfig[testConfig]()

	assert.Equal(t, expected, actual.FOO)

	// Restore previous state
	os.Setenv("FOO", initialState)
}

func TestOptionalDoesntExist(t *testing.T) {

	type testConfig struct {
		FOO string `env:"FOO" required:"false" default:"default value"`
	}

	// Save initial state and setup test env
	initialState := os.Getenv("FOO")
	os.Setenv("FOO", "")

	defer func() {
		r := recover()
		assert.Nil(t, r)
	}()

	actual := GenConfig[testConfig]()
	assert.Equal(t, "default value", actual.FOO)

	// Restore previous state
	os.Setenv("FOO", initialState)
}

func TestRequiredDoesntExistPanic(t *testing.T) {

	type testConfig struct {
		FOO string `env:"FOO" required:"true" default:"default value"`
	}

	// Save initial state and setup test env
	initialState := os.Getenv("FOO")
	os.Setenv("FOO", "")

	defer func() {
		r := recover()
		assert.NotNil(t, r)
		os.Setenv("FOO", initialState)
	}()

	GenConfig[testConfig]()

	// Restore previous state
	os.Setenv("FOO", initialState)
}

func genToken(length int) string {
	t := make([]byte, length)
	rand.Read(t)
	return base32.StdEncoding.EncodeToString(t)
}
