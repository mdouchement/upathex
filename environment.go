package upathex

import (
	"fmt"
	"os"
)

// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables. References to undefined
// variables left as there are.
func ExpandEnv(path string) string {
	return ExpandEnvWithCustom(path, nil)
}

// ExpandEnvWithCustom replaces ${var} or $var in the string according to the values
// of the given local map orof the current environment variables. References to undefined
// variables left as there are.
func ExpandEnvWithCustom(path string, local map[string]string) string {
	path = os.Expand(path, func(k string) string {
		if local != nil {
			if e, ok := local[k]; ok {
				return e
			}
		}

		if e := os.Getenv(k); e != "" {
			return e
		}

		return fmt.Sprintf("${%s}", k)
	})

	return path
}
