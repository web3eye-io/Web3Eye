package netutils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/web3eye-io/Web3Eye/config"
)

const (
	defaultHTTPTimeout = 30

// defaultHTTPKeepAlive           = 600
// defaultHTTPMaxIdleConns        = 100
// defaultHTTPMaxIdleConnsPerHost = 100
)

var (
	DefaultHTTPClient = newHTTPClientForRPC()
	// defaultMetricsHandler       = metricsHandler{}
	maxRespBodySize int64 = 1024 * 1024 * 1024
	errDataTooLarge       = errors.New("data too large")
)

// metricsHandler traces RPC records that get logged by the RPC client
// type metricsHandler struct{}

// ErrHTTP represents an error returned from an HTTP request
type ErrHTTP struct {
	URL    string
	Status int
}

func (h ErrHTTP) Error() string {
	return fmt.Sprintf("HTTP Error Status - %d | URL - %s", h.Status, h.URL)
}

func parseContentType(contentType string) string {
	contentType = strings.TrimSpace(contentType)
	whereCharset := strings.IndexByte(contentType, ';')
	if whereCharset != -1 {
		contentType = contentType[:whereCharset]
	}
	return contentType
}

func parseContentLength(contentLength string) (int64, error) {
	if contentLength != "" {
		contentLengthInt, err := strconv.Atoi(contentLength)
		if err != nil {
			return 0, err
		}
		return int64(contentLengthInt), nil
	}
	return 0, nil
}

func getContentHeaders(ctx context.Context, url string) (contentType string, contentLength int64, err error) {
	// Check if server supports HEAD
	if headers, err := getHeaders(ctx, "HEAD", url); err == nil {
		contentType = parseContentType(headers.Get("Content-Type"))
		contentLength, err = parseContentLength(headers.Get("Content-Length"))
		return contentType, contentLength, err
	}

	// Otherwise try GET
	headers, err := getHeaders(ctx, "GET", url)
	if err == nil {
		contentType = parseContentType(headers.Get("Content-Type"))
		contentLength, err = parseContentLength(headers.Get("Content-Length"))
		return contentType, contentLength, err
	}

	return contentType, contentLength, err
}

func getHeaders(ctx context.Context, method, url string) (http.Header, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := DefaultHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 399 || resp.StatusCode < 200 {
		return nil, ErrHTTP{Status: resp.StatusCode, URL: url}
	}

	defer resp.Body.Close()
	return resp.Header, nil
}

// newHTTPClientForRPC returns an http.Client configured with default settings intended for RPC calls.
func newHTTPClientForRPC() *http.Client {
	// get x509 cert pool
	// pool, err := x509.SystemCertPool()
	// if err != nil {
	// 	panic(err)
	// }

	// // walk every file in the tls directory and add them to the cert pool
	// err = filepath.WalkDir("_deploy/root-certs", func(path string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if d.IsDir() {
	// 		return nil
	// 	}
	// 	bs, err := os.ReadFile(path)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// append cert to pool
	// 	ok := pool.AppendCertsFromPEM(bs)
	// 	if !ok {
	// 		return fmt.Errorf("failed to append cert to pool")
	// 	}
	// 	return nil
	// })

	// // TODO: process the err
	// if err != nil {
	// 	return nil
	// }

	return &http.Client{
		Timeout: time.Second * defaultHTTPTimeout,
	}
}

func GetHTTPHeaders(ctx context.Context, url string) (contentType string, contentLength int64, err error) {
	return getContentHeaders(ctx, url)
}

// GetIPFSHeaders returns the headers for the given IPFS hash
func GetIPFSHeaders(ctx context.Context, path string) (contentType string, contentLength int64, err error) {
	url := fmt.Sprintf("%s/ipfs/%s", config.GetConfig().IPFS.HTTPGateway, path)
	return getContentHeaders(ctx, url)
}

func GetIPFSData(ctx context.Context, path string) ([]byte, error) {
	url := fmt.Sprintf("%s/ipfs/%s", config.GetConfig().IPFS.HTTPGateway, path)
	req, err := http.NewRequestWithContext(ctx, "GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	resp, err := DefaultHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 399 || resp.StatusCode < 200 {
		return nil, ErrHTTP{Status: resp.StatusCode, URL: url}
	}
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	if _, err := io.CopyN(buf, resp.Body, maxRespBodySize); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	extra := make([]byte, 1)
	if n, _ := io.ReadFull(resp.Body, extra); n > 0 {
		return nil, errDataTooLarge
	}

	return buf.Bytes(), nil
}

// GetURIPath takes a uri in any form and returns just the path
func GetURIPath(initial string, withoutQuery bool) string {
	var path string

	path = strings.TrimSpace(initial)
	if strings.HasPrefix(initial, "http") {
		path = strings.TrimPrefix(path, "https://")
		path = strings.TrimPrefix(path, "http://")
		indexOfPath := strings.Index(path, "/")
		if indexOfPath > 0 {
			path = path[indexOfPath:]
		}
	} else if strings.HasPrefix(initial, "ipfs://") {
		path = strings.ReplaceAll(initial, "ipfs://", "")
	} else if strings.HasPrefix(initial, "arweave://") || strings.HasPrefix(initial, "ar://") {
		path = strings.ReplaceAll(initial, "arweave://", "")
		path = strings.ReplaceAll(path, "ar://", "")
	}
	path = strings.ReplaceAll(path, "ipfs/", "")
	path = strings.TrimPrefix(path, "/")
	if withoutQuery {
		path = strings.Split(path, "?")[0]
		path = strings.TrimSuffix(path, "/")
	}

	return path
}

func GetArweaveDataHTTPReader(ctx context.Context, id string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://arweave.net/%s", id), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error getting data: %s", err.Error())
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting data: %s", err.Error())
	}
	return resp.Body, nil
}

func GetArweaveDataHTTP(ctx context.Context, id string) ([]byte, error) {
	resp, err := GetArweaveDataHTTPReader(ctx, id)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	data, err := io.ReadAll(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading data: %s", err.Error())
	}
	return data, nil
}
