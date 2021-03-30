package pkg

import "strings"

var (
	starts = map[string]string {
		"feat": "new feature",
		"fix": "fix bug",
		"docs": "documentation",
		"style": "code style without business change",
		"refactor": "refactor code",
		"test": "test code",
		"chore": "change for build tools",
	}
)

// CheckPrefix check the commit msg's prefix.
func CheckPrefix(args string) bool {
	for key, _ := range starts {
		if strings.HasPrefix(args, key) {
			return true
		}
	}
	return false
}
