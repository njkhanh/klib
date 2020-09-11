package cloudflare

import (
	"net/http"
	"strings"
)

var (
	client = &http.Client{}
)

func IsRestricted(url string) bool {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(req)
	defer res.Body.Close()
	return isRestricted(res)
}

func isRestricted(r *http.Response) bool {
	if r.StatusCode == 503 && strings.Contains(r.Header.Get("Server"), "cloudflare") {
		return true
	}
	return false
}
