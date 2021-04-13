package utils

import (
	"encoding/base64"
	"io"
	"os"
	"strings"
)

//SaveBase64StringToFile function to save base64 data to file
func SaveBase64StringToFile(path string, fileNameWithoutType string, encodedBase64 string) (string, int64, error) {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	_fileType := encodedBase64[0:strings.Index(encodedBase64, ";")]
	fileType := _fileType[strings.Index(_fileType, "/")+1:]
	filePath := path + "/" + fileNameWithoutType + "." + fileType
	encodedFileData := encodedBase64[strings.Index(encodedBase64, ",")+1:]

	fileBase64Decoded := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encodedFileData))

	storedFile, err := os.Create(filePath)
	defer storedFile.Close()

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", 0, err
	}
	fileSize := fileInfo.Size()

	if err != nil {
		return "", 0, err
	}

	_, err = io.Copy(storedFile, fileBase64Decoded)
	if err != nil {
		return "", 0, err
	}

	return filePath, fileSize, nil
}
