package utils

import (
	"fmt"
	"sort"
	"strings"
)

func Env(m map[string]interface{}) string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for _, key := range keys {
		fmt.Fprintf(&builder, "%s=%v\n", key, m[key])
	}
	return strings.TrimRight(builder.String(), "\n")
}

