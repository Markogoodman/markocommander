package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderscore(s string) string {
	var res []rune
	for i, r := range s {
		if i == 0 {
			res = append(res, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			res = append(res, '_')
		}
		res = append(res, unicode.ToLower(r))
	}
	return string(res)
}
