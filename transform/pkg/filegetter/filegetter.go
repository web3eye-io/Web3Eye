package filegetter

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
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

func Base64SVGToPng(base64data, dirPath, fileName string) (pngPath *string, err error) {
	svgPath := fmt.Sprintf("%v/%v.%v", dirPath, fileName, "svg")
	err = Base64SVGToSVG(base64data, svgPath)
	if err != nil {
		return nil, err
	}
	defer os.Remove(svgPath)

	_pngPath := fmt.Sprintf("%v/%v.%v", dirPath, fileName, "png")
	err = SVGToPng(svgPath, _pngPath)
	if err != nil {
		return nil, err
	}

	return &_pngPath, nil
}

func SVGToPng(svgPath, pngPath string) error {
	width, height := 256.0, 256.0
	if w, h, err := ReadSVGSizeByViewBox(svgPath); err == nil {
		width, height = w, h
	}

	chrome := svg2png.NewChrome().SetHeight(int(height)).SetWith(int(width))
	if err := chrome.Screenshoot(svgPath, pngPath); err != nil {
		logrus.Panic(err)
	}
	return nil
}

func Base64SVGToSVG(base64data string, svgPath string) (err error) {
	if !strings.HasPrefix(base64data, Base64SVGPrefix) {
		return errors.New("url format is not base64 svg")
	}

	noHeadData := strings.TrimPrefix(base64data, Base64SVGPrefix)
	rawData, err := base64.StdEncoding.DecodeString(noHeadData)
	if err != nil {
		return err
	}

	// create file to recive file stream
	out, err := os.Create(svgPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(rawData)
	if err != nil {
		return err
	}

	return nil
}

func ReadSVGSizeByViewBox(filePath string) (float64, float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	r, err := regexp.Compile("viewBox=\"([^\"]*)\"")
	if err != nil {
		return 0, 0, err
	}

	maxW, maxH := 0.0, 0.0
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if items := r.FindAllString(string(lineBytes), -1); len(items) > 0 {
			// find all string like 'viewBox="0 0 256.1 256.3"'
			// then parse it to width and height
			for _, v := range items {
				v = strings.Trim(v, "viewBox=")
				v = strings.Trim(v, "\"")
				nums := strings.Split(v, " ")
				if len(nums) < 4 {
					continue
				}
				width, err := strconv.ParseFloat(nums[2], 32)
				if err != nil {
					continue
				}
				height, err := strconv.ParseFloat(nums[3], 32)
				if err != nil {
					continue
				}
				if maxW < width {
					maxW = width
				}
				if maxH < height {
					maxH = height
				}
			}
		}
	}
	return maxW, maxH, nil
}

func ReadSVGSizeByWH(filePath string) (float64, float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(" width=\"([^\"]*)\"| height=\"([^\"]*)\"")
	if err != nil {
		return 0, 0, err
	}

	maxW, maxH := 0.0, 0.0
	for scanner.Scan() {
		if items := r.FindAllString(scanner.Text(), -1); len(items) > 0 {
			// find all string like 'width="0 0 256.1 256.3"'
			// then parse it to width and height
			for _, v := range items {
				v = strings.ReplaceAll(v, "\"", "")
				v = strings.TrimSpace(v)
				width, err := strconv.ParseFloat(strings.TrimPrefix(v, "width="), 32)
				if err == nil && maxW < width {
					maxW = width
				}

				height, err := strconv.ParseFloat(strings.TrimPrefix(v, "height="), 32)
				if err == nil && maxH < height {
					maxH = height
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return maxW, maxH, nil
}
