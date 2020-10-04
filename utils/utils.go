package utils

import (
	"strings"
)

func IsNotEmptyStr(s string) bool {
	return !(strings.TrimSpace(s) == "")
}
