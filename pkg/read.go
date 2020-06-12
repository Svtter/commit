package pkg

import (
	"os"
	"strings"
)

// LoadArgs read args from command line
func LoadArgs() string {
	argsWithoutProg := os.Args[1:]
	combineString := strings.Join(argsWithoutProg, " ")
	return combineString
}
