package upload

import (
	"fmt"
	upload "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"
	"github.com/google/uuid"
)

func (s Service) Upload(file upload.File) (shareURL string, err error) {
	fileName := file.Filename

	if fileName == "" {
		fileName = fmt.Sprintf("%s.%s", uuid.New().String(), file.Ext)
	}

	if err = s.createEmptyFile(file, fileName); err != nil {
		return shareURL, err
	}

	if err = s.fillEmptyFile(file, fileName); err != nil {
		return shareURL, err
	}

	return fmt.Sprintf("%s/%s/%s?inline=true", file.BaseURL, file.ShareURL, fileName), nil
}
