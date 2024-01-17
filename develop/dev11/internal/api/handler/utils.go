package handler

import (
	"net/url"
	"strings"
)

func formToStringMap(form url.Values) map[string]string {
	result := make(map[string]string)

	for k, v := range form {
		switch len(form[k]) {
		case 0:
			result[k] = ""
		case 1:
			result[k] = v[0]
		default:
			result[k] = strings.Join(v, " ")
		}
	}

	return result
}
