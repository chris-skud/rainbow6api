package rainbow6api

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *API) request(url string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", "Ubi_v1 t="+a.Session.Ticket)
	r.Header.Add("Ubi-AppId", a.appID)
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore issue with self-signed cert
		},
	}

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading error body %s, status code %d", err.Error(), resp.StatusCode)
		}

		var ubiErr ErrLogin
		err = json.Unmarshal(b, &ubiErr)
		if err != nil {
			return nil, fmt.Errorf("error Unmarshalling error body %s, status code %d", err.Error(), resp.StatusCode)
		}

		return nil, errors.New("ubisoft request error with status code " + strconv.Itoa(resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}
