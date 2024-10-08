package contentreader

import (
	"errors"
	"path/filepath"
	"regexp"
)

func ReadText(fileName string, document []byte) (content string, err error) {
	ext := getFileExtension(fileName)
	switch ext {
	case ".docx":
		content, err = readDocx(document)
		return
	case ".pdf":
		content, err = readPdf(document)
		return
	}

	return "", errors.New("unsupported file type")
}

func getFileExtension(fileName string) string {
	ext := filepath.Ext(fileName)
	regex := regexp.MustCompile(`\?.+`)
	return regex.ReplaceAllString(ext, "")
}
