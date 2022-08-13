package upathex_test

import (
	"os"
	"testing"

	"github.com/mdouchement/upathex"
	"github.com/stretchr/testify/assert"
)

func TestExpandEnv(t *testing.T) {
	_, err := os.Stat("/.dockerenv")
	if os.IsNotExist(err) {
		assert.FailNow(t, "The test must be runned in a container as root")
	}

	//

	tests := []struct {
		raw      string
		expected string
		env      map[string]string
	}{
		{
			raw:      "$HOME",
			expected: "/root",
		},
		{
			raw:      "${HOME}",
			expected: "/root",
		},
		//
		//
		//
		{
			raw:      "/usr/$CUSTOM_ENV/bin",
			expected: "/usr/local/bin",
			env: map[string]string{
				"CUSTOM_ENV": "local",
			},
		},
	}

	for _, test := range tests {
		for k, v := range test.env {
			os.Setenv(k, v)
		}

		path := upathex.ExpandEnv(test.raw)
		assert.Equal(t, test.expected, path, test.raw)
	}
}

func TestExpandEnvWithCustom(t *testing.T) {
	_, err := os.Stat("/.dockerenv")
	if os.IsNotExist(err) {
		assert.FailNow(t, "The test must be runned in a container as root")
	}

	//

	tests := []struct {
		raw      string
		expected string
		env      map[string]string
	}{
		{
			raw:      "$HOME",
			expected: "/root",
		},
		{
			raw:      "${HOME}",
			expected: "/root",
		},
		//
		//
		//
		{
			raw:      "/usr/$CUSTOM_LOCAL/bin",
			expected: "/usr/local/bin",
			env: map[string]string{
				"CUSTOM_LOCAL": "local",
			},
		},
	}

	for _, test := range tests {
		path := upathex.ExpandEnvWithCustom(test.raw, test.env)
		assert.Equal(t, test.expected, path, test.raw)
	}
}
