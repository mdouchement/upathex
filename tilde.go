package upathex

import (
	"os/user"
	"path/filepath"
	"strings"
)

const separator = string(filepath.Separator)

// ExpandTilde replaces `~/' by the current user home directory.
// It also replaces `~mdouchement/' by the mdouchement home directory.
func ExpandTilde(path string) (string, error) {
	// Current user.
	if path == "~" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return usr.HomeDir, nil
	}

	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		path = strings.Replace(path, "~", usr.HomeDir, 1)
		path = strings.TrimRight(path, separator)
		return path, nil
	}

	// Another user (e.g. ~mdouchement/).
	if strings.HasPrefix(path, "~") {
		username, path, found := strings.Cut(path[1:], separator)

		usr, err := user.Lookup(username)
		if err != nil {
			return "", err
		}

		if !found {
			return usr.HomeDir, nil
		}

		return filepath.Join(usr.HomeDir, path), nil
	}

	return path, nil
}
