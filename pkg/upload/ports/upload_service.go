package upload

import "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"

type Service interface {
	Upload(file upload.File) (shareURL string, err error)
}
