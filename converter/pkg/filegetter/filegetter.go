package filegetter

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/canhlinh/svg2png"
	"github.com/sirupsen/logrus"
)

func DownloadAreaveFile(url string, dirPath string, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, "ar://") {
		return nil, errors.New("url format is not areave")
	}

	noHeadUrl := strings.TrimPrefix(url, "ar://")
	httpUrl := fmt.Sprintf("%v/%v", "https://arweave.net", noHeadUrl)
	return DownloadHttpFile(httpUrl, dirPath, fileName)
}

func DownloadIPFSFile(url string, dirPath string, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, "ipfs://") {
		return nil, errors.New("url format is not areave")
	}

	noHeadUrl := strings.TrimPrefix(url, "ipfs://")
	httpUrl := fmt.Sprintf("%v/%v", "https://ipfs.io/ipfs", noHeadUrl)
	return DownloadHttpFile(httpUrl, dirPath, fileName)
}

func DownloadHttpFile(url string, dirPath string, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return nil, errors.New("url format is not http")
	}

	// get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get content-type
	typeTree := strings.Split(resp.Header.Get("Content-Type"), "/")

	var formatDetect []byte
	// if not get right current format,detect content type
	if typeTree[len(typeTree)-1] == "octet-stream" {
		formatDetect = make([]byte, 512)
		resp.Body.Read(formatDetect)

		contentType := http.DetectContentType(formatDetect)
		typeTree = strings.Split(contentType, "/")
	}

	// create file to recive file stream
	filePath := fmt.Sprintf("%v/%v.%v", dirPath, fileName, typeTree[len(typeTree)-1])
	out, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	defer out.Close()

	_, err = out.Write(formatDetect)
	if err != nil {
		return nil, err
	}

	// from resp body get file stream
	_, err = io.Copy(out, resp.Body)
	return &filePath, err
}

func Base64ToSVG(base64data string, dirPath string, fileName string) (path *string, err error) {
	base64prefix := "data:image/svg+xml;base64,"
	if !strings.HasPrefix(base64data, base64prefix) {
		return nil, errors.New("url format is not base64 svg")
	}

	noHeadData := strings.TrimPrefix(base64data, base64prefix)
	rawData, err := base64.StdEncoding.DecodeString(noHeadData)
	if err != nil {
		return nil, err
	}

	// create file to recive file stream
	svgPath := fmt.Sprintf("%v/%v.%v", dirPath, fileName, "svg")
	out, err := os.Create(svgPath)
	if err != nil {
		return nil, err
	}

	defer out.Close()
	defer os.Remove(svgPath)

	_, err = out.Write(rawData)
	if err != nil {
		return nil, err
	}

	chrome := svg2png.NewChrome().SetHeight(600).SetWith(600)
	pngPath := fmt.Sprintf("%v/%v.%v", dirPath, fileName, "png")
	if err := chrome.Screenshoot(svgPath, pngPath); err != nil {
		logrus.Panic(err)
	}

	return &pngPath, err
}
