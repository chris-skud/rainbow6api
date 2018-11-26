package rainbow6api

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const loginURL = "https://connect.ubi.com/ubiservices/v2/profiles/sessions"

func (a *API) login() error {
	r, err := http.NewRequest(http.MethodPost, loginURL, bytes.NewReader([]byte(`{"rememberMe":true}`)))
	if err != nil {
		return err
	}

	authStr := fmt.Sprintf("%s:%s", a.email, a.password)
	b64Creds := base64.StdEncoding.EncodeToString([]byte(authStr))

	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("Accept", "*/*")
	r.Header.Add("Ubi-AppId", a.appID)
	r.Header.Add("Ubi-RequestedPlatformType", a.platform)
	r.Header.Add("Authorization", "Basic "+b64Creds)
	r.Header.Add("X-Requested-With", "XMLHttpRequest")
	r.Header.Add("Referer", "https://connect.ubi.com/Default/Login?appId="+a.appID+"&lang=en-US&nextUrl=https%3A%2F%2Fclub.ubisoft.com%2Flogged-in.html%3Flocale%3Den-US")
	r.Header.Add("Accept-Language", "en-US")
	r.Header.Add("Accept-Encoding", "deflate, br")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	r.Header.Add("Host", "connect.ubi.com")
	r.Header.Add("Content-Length", "19")
	r.Header.Add("Cache-Control", "no-cache")

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore issue with self-signed cert
		},
	}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading error body %s, status code %d", err.Error(), resp.StatusCode)
		}

		var ubiErr ErrLogin
		err = json.Unmarshal(b, &ubiErr)
		if err != nil {
			return fmt.Errorf("error Unmarshalling error body %s, status code %d", err.Error(), resp.StatusCode)
		}

		return errors.New("login request failed: " + ubiErr.Error())
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var session Session
	err = json.Unmarshal(b, &session)
	if err != nil {
		return err
	}

	a.Session = &session
	return nil
}

func (a *API) checkSession() error {
	// if the ses is set, and the session has not expired, return ok
	if a.Session != nil && a.Session.Expiration.Before(time.Now().Add(-1*time.Second)) {
		return nil
	}

	return a.login()
}
