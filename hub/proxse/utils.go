package main

import (
	"strings"
	"errors"
)

func URLGetSessId(URL string) (string, error) {
	if !strings.HasPrefix(URL, "/session/") {
		return "", errors.New("Not a /session/<id> url")
	}
	s ,_ := strings.CutPrefix(URL, "/session/")
	s, _, _ = strings.Cut(s, "/")
	return s, nil
}
