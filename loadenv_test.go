package loadenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testEnv = `FOO=fooVal
BAR=
BAZ=bazVal`

func TestLoadEnv(t *testing.T) {
	// Store current environment and setup testing environment
	var err error
	initialState := map[string]string{}
	usedVars := []string{"FOO", "BAR", "BAZ"}
	for _, val := range usedVars {
		initialState[val] = os.Getenv(val)
		os.Setenv(val, "")
	}
	fName := genToken(15)
	f, err := os.Create("." + fName)
	if err != nil {
		t.Fatalf("failed to create environment file for testing: %s", err)
	}
	_, err = f.WriteString(testEnv)
	if err != nil {
		t.Fatalf("failed to write .env file for testing: %s", err)
	}
	f.Close()

	err = LoadEnv("." + fName)
	assert.Nil(t, err)

	assert.Equal(t, "fooVal", os.Getenv("FOO"))
	assert.Equal(t, "", os.Getenv("BAR"))
	assert.Equal(t, "bazVal", os.Getenv("BAZ"))

	// Cleanup
	for _, val := range usedVars {
		os.Setenv(val, initialState[val])
	}
	os.Remove("." + fName)
}
