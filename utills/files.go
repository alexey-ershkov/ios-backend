package utills

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	uuid "github.com/nu7hatch/gouuid"
	"ios-backend/configs"
)

func SaveFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {

	defer file.Close()

	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	uString := u.String()
	folderName := []rune(uString)[:3]
	separatedFilename := strings.Split(header.Filename, ".")
	if len(separatedFilename) <= 1 {
		err := errors.New("bad filename")
		return "", err
	}
	fileType := separatedFilename[len(separatedFilename)-1]

	path := fmt.Sprintf("%s/%s/%s", configs.MEDIA_FOLDER, folder, string(folderName))
	filename := fmt.Sprintf("%s.%s", uString, fileType)

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", nil
	}

	fullFilename := fmt.Sprintf("%s/%s", path, filename)

	f, err := os.OpenFile(fullFilename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	return fullFilename, err
}
