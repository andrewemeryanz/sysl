package importer

import (
	"net/url"
	"strings"
)

func getDescription(d string) string {
	if d == "" {
		return "No description."
	}
	return d
}

func quote(s string) string {
	if s == "" {
		return ""
	}
	return `"` + s + `"`
}

func isBuiltInType(item Type) bool {
	switch item.(type) {
	case *SyslBuiltIn, *ImportedBuiltInAlias:
		return true
	}
	return false
}

func isUnionType(item Type) bool {
	if _, ok := item.(*Union); ok {
		return true
	}
	return false
}

func isExternalAlias(item Type) bool {
	switch item.(type) {
	case *ExternalAlias, *Array, *Enum, *Alias:
		return true
	}
	return false
}

func getSyslTypeName(item Type) string {
	switch t := item.(type) {
	case *Array:
		return "sequence of " + getSyslTypeName(t.Items)
	case *Enum:
		return item.Name()
	case *ImportedBuiltInAlias:
		return getSyslTypeName(t.Target)
	case *Alias:
		return t.Name()
	case *Union:
		return t.Name()
	}
	if isExternalAlias(item) {
		return "EXTERNAL_" + item.Name()
	}
	return item.Name()
}

func spaceSeparate(items ...string) string {
	var t []string
	for _, i := range items {
		if i != "" {
			t = append(t, i)
		}
	}
	return strings.Join(t, " ")
}

func getSyslSafeEndpoint(endpoint string) string {
	// url.PathEscape does not escape '.'
	charsToKeep := map[string]string{
		`%2F`: "/",
		`%7B`: "{",
		`%7D`: "}",
		`%3D`: "=",
		`%3F`: "?",
		`%26`: "&",
	}
	endpoint = url.PathEscape(endpoint)
	endpoint = strings.ReplaceAll(endpoint, ".", `%2E`)
	for hex, realChar := range charsToKeep {
		endpoint = strings.ReplaceAll(endpoint, hex, realChar)
	}
	return endpoint
}
