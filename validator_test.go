package gladiator

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/displae/govalidator"
	"github.com/displae/testify/assert"
)

// Test input validation of displae username
// Displae username must only contain alphanumeric,
// dot and underscore characters. The maximum number
// of displae username characters are 25
func TestValidateUsername(t *testing.T) {
	usernameTemplate := "displae{{char}}username"
	type InputUsername struct {
		Username string
		Expected bool
	}

	var inputUsernames []InputUsername

	for asciiCode := 0; asciiCode < 128; asciiCode++ {
		username := strings.Replace(usernameTemplate, "{{char}}", string(asciiCode), 1)
		expected := false
		switch {
		case govalidator.IsAlphanumeric(username):
			expected = true
		case strings.Contains(username, "_"):
			expected = true
		case strings.Contains(username, "."):
			expected = true
		default:
			expected = false
		}

		inputUsernames = append(inputUsernames, InputUsername{Username: username, Expected: expected})
	}

	for _, inputUsername := range inputUsernames {
		actual := ValidateDisplaeUsername(inputUsername.Username)
		assert.Equal(t, inputUsername.Expected, actual, "Username considered not valid but valid or considered valid but not valid.")
	}
}

// Test validation of uncommon characters input such as chinese
// characters, strange characters and other unknown characters.
func TestValidateDisplaeUsernameUnknownChars(t *testing.T) {
	usernameTemplate := "displae{{char}}username"
	numNegative := rand.Int63() * -1
	numGt128 := rand.Int63n(128) + int64(rand.Int())

	inputs := []int64{numGt128, numNegative}
	for _, input := range inputs {
		username := strings.Replace(usernameTemplate, "{{char}}", string(input), 1)
		actual := ValidateDisplaeUsername(username)
		assert.Equal(t, false, actual, "Username must be invalid")
	}
}
