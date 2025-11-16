package utils

import (
	"strings"
	"strconv"
	"errors"
)

// Extracts session id from url
func URLGetSessId(URL string) (string, error) {
	if !strings.HasPrefix(URL, "/session/") {
		return "", errors.New("not a /session/<id> url")
	}
	s, _ := strings.CutPrefix(URL, "/session/")
	s, _, _ = strings.Cut(s, "/")
	return s, nil
}

// Extracts test id from url
func URLGetTestId(URL string) (int64, error) {
	if !strings.HasPrefix(URL, "/thex/test/") {
		return 0, errors.New("not a /thex/test/<id> url")
	}
	s, _ := strings.CutPrefix(URL, "/session/")
	s, _, _ = strings.Cut(s, "/")
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}