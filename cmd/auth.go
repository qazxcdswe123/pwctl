package cmd

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func getJWTToken(ctx context.Context, username, password, url string) (string, error) {
	url = url + "/login"
	body := `{"user":"` + username + `","password":"` + password + `"}`
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	respString := string(respBody)
	return respString, nil
}
