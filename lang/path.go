package lang

import (
	"net/url"
	"strings"
)

func Path()string {
	rawurl := "http://DOMAIN/bfs/${bucketname}/${filename}"
	var urlValue *url.URL
	var err error
	if urlValue, err = url.ParseRequestURI(rawurl); err != nil {
		return ""
	}

	s := strings.Split(urlValue.Path, "/")
	if len(s) != 4 {
		return ""
	}

	return s[2] + " | " + s[3]
}