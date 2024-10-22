package upload

import (
	"fmt"
	uploadDomain "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"
	"net/http"
)

func (s Service) CreateEmptyFile(baseData uploadDomain.BaseData) (shareURL string, err error) {
	var client = &http.Client{}

	var url = fmt.Sprintf("%s/tus/%s/%s?override=false", baseData.BaseURL, baseData.Folder, baseData.Filename)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("X-Auth", baseData.Token)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	shareURL = fmt.Sprintf("%s/%s/%s?inline=true", baseData.BaseURL, baseData.ShareURL, baseData.Filename)

	return shareURL, nil
}
