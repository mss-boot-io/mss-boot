package reader

import (
	"os"
	"regexp"
)

// ReplaceEnvVars replace env vars
func ReplaceEnvVars(raw []byte) ([]byte, error) {
	re := regexp.MustCompile(`\$\{([A-Za-z0-9_]+)\}`)
	if re.Match(raw) {
		dataS := string(raw)
		res := re.ReplaceAllStringFunc(dataS, replaceEnvVars)
		return []byte(res), nil
	}
	return raw, nil
}

func replaceEnvVars(element string) string {
	v := element[2 : len(element)-1]
	el := os.Getenv(v)
	return el
}
