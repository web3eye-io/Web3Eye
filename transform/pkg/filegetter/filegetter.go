package filegetter

import (
	"bufio"
	"bytes"
	"context"
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
	ArURLHead        = "ar://"
	ArURLHTTPGateway = "https://arweave.net"
	IPFSUrlHead      = "ipfs://"
	IPFSHttpGateway  = "https://ipfs.io/ipfs"
	HTTPUrlHead      = "http://"
	HTTPSUrlHead     = "https://"
	Base64SVGPrefix  = "data:image/svg+xml;base64,"
	whBitsize        = 32
)

func GetFileFromURL(ctx context.Context, url, dirPath, fileName string) (path *string, err error) {
	switch {
	case strings.HasPrefix(url, ArURLHead):
		return DownloadAreaveFile(ctx, url, dirPath, fileName)
	case strings.HasPrefix(url, IPFSUrlHead):
		return DownloadIPFSFile(ctx, url, dirPath, fileName)
	case strings.HasPrefix(url, Base64SVGPrefix):
		return Base64SVGToPng(url, dirPath, fileName)
	default:
		return DownloadHTTPFile(ctx, url, dirPath, fileName)
	}
}

func DownloadAreaveFile(ctx context.Context, url, dirPath, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, ArURLHead) {
		return nil, errors.New("url format is not areave")
	}

	noHeadURL := strings.TrimPrefix(url, ArURLHead)
	httpURL := fmt.Sprintf("%v/%v", ArURLHTTPGateway, noHeadURL)
	return DownloadHTTPFile(ctx, httpURL, dirPath, fileName)
}

func DownloadIPFSFile(ctx context.Context, url, dirPath, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, IPFSUrlHead) {
		return nil, errors.New("url format is not areave")
	}

	noHeadURL := strings.TrimPrefix(url, IPFSUrlHead)
	httpURL := fmt.Sprintf("%v/%v", IPFSHttpGateway, noHeadURL)
	return DownloadHTTPFile(ctx, httpURL, dirPath, fileName)
}

func DownloadHTTPFile(ctx context.Context, url, dirPath, fileName string) (path *string, err error) {
	if !strings.HasPrefix(url, HTTPUrlHead) && !strings.HasPrefix(url, HTTPSUrlHead) {
		return nil, errors.New("url format is not http")
	}

	body := bytes.NewBuffer([]byte{})
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get content-type
	typeTree := strings.Split(resp.Header.Get("Content-Type"), "/")

	var formatDetect []byte
	// if not get right current format,detect content type
	if typeTree[len(typeTree)-1] == "octet-stream" {
		formatDataLen := 512
		formatDetect = make([]byte, formatDataLen)
		_, err = resp.Body.Read(formatDetect)
		if err != nil {
			return nil, err
		}

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

func Base64SVGToSVG(base64data, svgPath string) (err error) {
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

func ReadSVGSizeByViewBox(filePath string) (w, h float64, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	r := regexp.MustCompile("viewBox=\"([^\"]*)\"")
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

				viewBoxItemsLen := 4
				wIndex := 2
				hIndex := 3
				if len(nums) < viewBoxItemsLen {
					continue
				}
				width, err := strconv.ParseFloat(nums[wIndex], whBitsize)
				if err != nil {
					continue
				}
				height, err := strconv.ParseFloat(nums[hIndex], whBitsize)
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

func ReadSVGSizeByWH(filePath string) (w, h float64, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(" width=\"([^\"]*)\"| height=\"([^\"]*)\"")
	maxW, maxH := 0.0, 0.0
	for scanner.Scan() {
		if items := r.FindAllString(scanner.Text(), -1); len(items) > 0 {
			// find all string like 'width="0 0 256.1 256.3"'
			// then parse it to width and height
			for _, v := range items {
				v = strings.ReplaceAll(v, "\"", "")
				v = strings.TrimSpace(v)
				width, err := strconv.ParseFloat(strings.TrimPrefix(v, "width="), whBitsize)
				if err == nil && maxW < width {
					maxW = width
				}

				height, err := strconv.ParseFloat(strings.TrimPrefix(v, "height="), whBitsize)
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
