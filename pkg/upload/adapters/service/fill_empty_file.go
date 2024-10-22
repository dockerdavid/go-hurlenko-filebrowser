package upload

import (
	"bytes"
	"fmt"
	uploadDomain "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"
	"io"
	"net/http"
)

func (s Service) FillEmptyFile(baseData uploadDomain.BaseData, file uploadDomain.File) error {
	client := &http.Client{}

	url := fmt.Sprintf("%s/tus/%s/%s?override=false", baseData.BaseURL, baseData.Folder, baseData.Filename)

	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, bytes.NewReader(file.File))

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url, buf)

	if err != nil {
		return err
	}

	req.Header.Add("X-Auth", baseData.Token)
	req.Header.Add("Content-Type", "application/offset+octet-stream")
	req.Header.Add("Content-Length", fmt.Sprintf("%d", buf.Len()))
	req.Header.Add("Upload-Offset", "0")

	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: baseData.Token,
	})

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
