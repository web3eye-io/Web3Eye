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

const (
	ArUrlHead        = "ar://"
	ArUrlHttpGateway = "https://arweave.net"
	IPFSUrlHead      = "ipfs://"
	IPFSHttpGateway  = "https://ipfs.io/ipfs"
	HTTPUrlHead      = "http://"
	HTTPSUrlHead     = "https://"
	Base64SVGPrefix  = "data:image/svg+xml;base64,"
)

func GetFileFromURL(url string, dirPath string, fileName string) (path *string, err error) {
	switch {
	case strings.HasPrefix(url, ArUrlHead):
		return DownloadAreaveFile(url, dirPath, fileName)
	case strings.HasPrefix(url, IPFSUrlHead):
		return DownloadIPFSFile(url, dirPath, fileName)
	case strings.HasPrefix(url, Base64SVGPrefix):
		return Base64SVGToPng(url, dirPath, fileName)
	default:
		return DownloadHttpFile(url, dirPath, fileName)
	}
}

func DownloadAreaveFile(url string, dirPath string, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, ArUrlHead) {
		return nil, errors.New("url format is not areave")
	}

	noHeadUrl := strings.TrimPrefix(url, ArUrlHead)
	httpUrl := fmt.Sprintf("%v/%v", ArUrlHttpGateway, noHeadUrl)
	return DownloadHttpFile(httpUrl, dirPath, fileName)
}

func DownloadIPFSFile(url string, dirPath string, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, IPFSUrlHead) {
		return nil, errors.New("url format is not areave")
	}

	noHeadUrl := strings.TrimPrefix(url, IPFSUrlHead)
	httpUrl := fmt.Sprintf("%v/%v", IPFSHttpGateway, noHeadUrl)
	return DownloadHttpFile(httpUrl, dirPath, fileName)
}

func DownloadHttpFile(url string, dirPath string, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, HTTPUrlHead) && !strings.HasPrefix(url, HTTPSUrlHead) {
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

func Base64SVGToPng(base64data string, dirPath string, fileName string) (path *string, err error) {
	if !strings.HasPrefix(base64data, Base64SVGPrefix) {
		return nil, errors.New("url format is not base64 svg")
	}

	noHeadData := strings.TrimPrefix(base64data, Base64SVGPrefix)
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
