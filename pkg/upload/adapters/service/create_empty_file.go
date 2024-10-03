package upload

import (
	"fmt"
	upload "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"
	"net/http"
)

func (s Service) createEmptyFile(file upload.File, fileName string) error {
	var client = &http.Client{}

	var url = fmt.Sprintf("%s/tus/%s/%s?override=false", file.BaseURL, file.Folder, fileName)

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("X-Auth", file.Token)

	var resp *http.Response

	resp, err = client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
