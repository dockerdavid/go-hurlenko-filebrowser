package auth

import auth "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/auth/domain"

type Service interface {
	Auth(credentials auth.Credentials) (token string, err error)
}
