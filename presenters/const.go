package presenters

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrIdNotFound     = errors.New("Not Found")
	FileExtentionLogo = "jpeg,jpg,png"
)

func GetFileExtensionFromUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	pos := strings.LastIndex(u.Path, ".")
	if pos == -1 {
		return "", nil
	}
	return u.Path[pos+1 : len(u.Path)], nil
}

func ContainString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
