package upload

import (
	"fmt"
	uploadDomain "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"
	"net/http"
)

func (s Service) DeleteFile(delete uploadDomain.Delete) (err error) {
	var client = &http.Client{}

	var url = fmt.Sprintf("%s/resources/%s/%s", delete.BaseURL, delete.Folder, delete.Filename)

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("X-Auth", delete.Token)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("file not found %s", url)
	}

	return nil
}
