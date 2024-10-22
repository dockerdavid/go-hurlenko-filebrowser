package upload

import (
	uploadDomain "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"
)

func (s Service) Upload(baseData uploadDomain.BaseData, file uploadDomain.File) (shareURL string, err error) {
	shareURL, err = s.CreateEmptyFile(baseData)

	if err != nil {
		return shareURL, err
	}

	if err = s.FillEmptyFile(baseData, file); err != nil {
		return shareURL, err
	}

	return shareURL, nil
}
