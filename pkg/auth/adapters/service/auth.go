package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dockerdavid/go-hurlenko-filebrowser/pkg/auth/domain"
	"io"
	"net/http"
)

func (s Service) Auth(credentials auth.Credentials) (token string, err error) {
	authJson, err := json.Marshal(credentials)

	if err != nil {
		return "", err
	}

	var client = &http.Client{}

	var url = fmt.Sprintf("%s/login", credentials.BaseURL)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(authJson))

	if err != nil {
		return "", err
	}

	var resp *http.Response

	resp, err = client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
