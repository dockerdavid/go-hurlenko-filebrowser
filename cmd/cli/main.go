package main

import (
	"fmt"
	authService "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/auth/adapters/service"
	authDomain "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/auth/domain"
	uploadService "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/adapters/service"
	uploadDomain "github.com/dockerdavid/go-hurlenko-filebrowser/pkg/upload/domain"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	credentials := authDomain.Credentials{
		Username:  os.Getenv("USERNAME"),
		Password:  os.Getenv("PASSWORD"),
		Recaptcha: "",
		BaseURL:   os.Getenv("BASE_URL"),
	}

	auth := authService.Service{}

	token, err := auth.Auth(credentials)

	if err != nil {
		return
	}

	file := uploadDomain.File{
		Token:    token,
		BaseURL:  os.Getenv("BASE_URL"),
		Folder:   os.Getenv("FOLDER"),
		ShareURL: os.Getenv("SHARE_URL"),
		File:     []byte("test"),
		Ext:      "txt",
	}

	upload := uploadService.Service{}

	shareURL, err := upload.Upload(file)
	if err != nil {
		return
	}

	fmt.Println(shareURL)
}
