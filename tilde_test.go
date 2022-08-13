package upathex_test

import (
	"os"
	"os/exec"
	"os/user"
	"testing"

	"github.com/mdouchement/upathex"
	"github.com/stretchr/testify/assert"
)

func TestExpandTilde(t *testing.T) {
	_, err := os.Stat("/.dockerenv")
	if os.IsNotExist(err) {
		assert.FailNow(t, "The test must be runned in a container as root")
	}

	//

	tests := []struct {
		raw      string
		expected string
		useradd  string
	}{
		{
			raw:      "~",
			expected: "/root",
		},
		{
			raw:      "~/",
			expected: "/root",
		},
		{
			raw:      "~/bin",
			expected: "/root/bin",
		},
		//
		//
		//
		{
			raw:      "~mdouchement",
			expected: "/home/mdouchement",
			useradd:  "mdouchement",
		},
		{
			raw:      "~mdouchement/",
			expected: "/home/mdouchement",
		},
		{
			raw:      "~mdouchement/bin",
			expected: "/home/mdouchement/bin",
		},
	}

	for _, test := range tests {
		if test.useradd != "" {
			usr, err := user.Lookup(test.useradd)
			if err != nil || usr == nil {
				err = exec.Command("adduser", test.useradd).Run()
				assert.NoError(t, err, test.raw)
			}
		}

		path, err := upathex.ExpandTilde(test.raw)
		assert.NoError(t, err, test.raw)
		assert.Equal(t, test.expected, path, test.raw)
	}
}
