package urls

import (
	"net/http"
	"regexp"
)

// Find takes in a string (i.e. a discord message)
// and returns a list of all the urls it has found in this
// string.
func Find(s string) []string {
	// Regular expression to match URLs starting with https
	re := regexp.MustCompile(`https://[^\s]+`)

	// Find all matching URLs
	return re.FindAllString(s, -1)
}

func Resolve(url string) (string, error) {
	client := &http.Client{
		// Automatically follow redirects
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	resp, err := client.Head(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// The final resolved URL after following redirects
	finalURL := resp.Request.URL.String()
	return finalURL, nil
}
