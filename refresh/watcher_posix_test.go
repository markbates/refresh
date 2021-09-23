// +build linux darwin

package refresh

import (
	"testing"
)

func TestIsIgnoredFolder(t *testing.T) {
	ignoredFolders := []string{
		"cmd/web/client",
		"vendor",
	}

	isIgnoredFolderTests := []struct {
		path          string
		ignoredFolder bool
	}{
		{"cmd/web", false},
		{"cmd/web/main.go", false},
		{"cmd/web/client", true},
		{"cmd/web/client/src", true},
		{"pkg", false},
		{"pkg/cmd/web/client", false},
		{".", false},
	}

	watcher := Watcher{
		Manager: &Manager{
			Configuration: &Configuration{
				IgnoredFolders: ignoredFolders,
			},
		},
	}

	for _, tc := range isIgnoredFolderTests {
		if watcher.isIgnoredFolder(tc.path) != tc.ignoredFolder {
			if tc.ignoredFolder {
				t.Errorf("expected path '%s' to be ignored", tc.path)
			} else {
				t.Errorf("expected path '%s' not to be ignored", tc.path)
			}
		}
	}
}
