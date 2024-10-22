package upload

import "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"

type Service interface {
	Upload(baseData upload.BaseData, file upload.File) (shareURL string, err error)
	FillEmptyFile(baseDate upload.BaseData, file upload.File) error
	CreateEmptyFile(baseData upload.BaseData) (shareURL string, err error)
	DeleteFile(delete upload.Delete) (err error)
}
