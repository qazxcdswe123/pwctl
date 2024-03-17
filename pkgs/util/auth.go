package util

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func ValidateHostname(str string) error {
	u, err := url.Parse(str)
	if err != nil {
		return err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("invalid url scheme, must be http or https")
	}

	return nil
}

// Token is the jwt token to be used for authentication
// Maybe we can store this in a file with expiration time
var Token string

// SetJWTToken sets the JWT token
func SetJWTToken(username, password, url string) error {
	var ctx = context.Background()
	loginURL := url + "/login"
	body := `{"user":"` + username + `","password":"` + password + `"}`
	req, err := http.NewRequestWithContext(ctx, "POST", loginURL, strings.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respString := string(respBody)
	Token = respString
	return nil
}
