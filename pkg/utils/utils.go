package utils

import (
	log "github.com/Sirupsen/logrus"
	"net/url"
	"os"
	"regexp"

	"github.com/docker/go-plugins-helpers/authorization"
)

func StringInSlice(haystack []string, needle string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

func GetURIInfo(req authorization.Request) (string, string) {
	reURI := regexp.MustCompile(`^/(v[0-9]+\.[0-9]+)(/.*)`)

	result := reURI.FindStringSubmatch(req.RequestURI)

	return result[1], result[2]
}

func GetURLParams(r string) url.Values {
	u, err := url.ParseRequestURI(r)
	if err != nil {
		log.Fatal(err)
	}

	return u.Query()
}

func FileExists(f string) bool {
	_, err := os.Lstat(f)
	if err != nil {
		return false
	}

	return true
}
