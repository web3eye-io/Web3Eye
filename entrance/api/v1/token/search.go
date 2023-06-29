package token

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	entranceproto "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/token"

	"github.com/web3eye-io/Web3Eye/common/servermux"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/entrance/resource"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	"github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/token"
	"google.golang.org/grpc"

	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

// 8mb

const (
	MaxUploadFileSize = 1 << 10 << 10 << 3
	UploadFileFeild   = "UploadFile"
	PageLimitFeild    = "PageLimit"
)

type Img2VectorResp struct {
	Vector  []float32 `json:"vector"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}

type SearchToken struct {
	nftmetaproto.Token
	Distance float32
}

func init() {
	mux := servermux.AppServerMux()
	mux.HandleFunc("/search/file", Search)

	pages, err := fs.Sub(resource.ResPages, "pages")
	if err != nil {
		log.Fatalf("failed to load pages: %v", err)
	}
	mux.Handle("/", http.FileServer(http.FS(pages)))
}

// nolint
func Search(w http.ResponseWriter, r *http.Request) {
	startT := time.Now()
	respBody := make(map[string]interface{})
	defer func() {
		_respBody, err := json.Marshal(respBody)
		if err != nil {
			respBody["msg"] = fmt.Sprintf("json marshal response body fail, %v", err)
		}
		_, err = w.Write(_respBody)
		if err != nil {
			respBody["msg"] = fmt.Sprintf("write response body fail, %v", err)
		}
	}()

	_pageLimit := r.FormValue(PageLimitFeild)
	pageLimit, err := strconv.ParseUint(_pageLimit, 10, 32)

	if err != nil {
		respBody["msg"] = fmt.Sprintf("failed to parse feild PageLimit %v, %v", _pageLimit, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// judge weather filesize exceed max-size
	err = r.ParseMultipartForm(MaxUploadFileSize)
	if err != nil {
		respBody["msg"] = fmt.Sprintf("read file failed %v, %v", MaxUploadFileSize, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inT := time.Now()
	logger.Sugar().Infof("check params %v ms", inT.UnixMilli()-startT.UnixMilli())

	// convert to vector
	vector, err := ImgReqConvertVector(r)
	if err != nil {
		respBody["msg"] = fmt.Sprintf("image convert fail, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inT = time.Now()
	logger.Sugar().Infof("finish convert to vector %v ms", inT.UnixMilli()-startT.UnixMilli())

	token.UseCloudProxyCC()
	resp, err := token.Search(context.Background(), &rankerproto.SearchTokenRequest{
		Vector: vector,
		Limit:  uint32(pageLimit),
	})
	if err != nil {
		respBody["msg"] = fmt.Sprintf("search fail, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inT = time.Now()
	logger.Sugar().Infof("finish query id %v ms", inT.UnixMilli()-startT.UnixMilli())

	respBody["Msg"] = fmt.Sprintf("have %v infos", len(resp.Infos))
	respBody["Infos"] = resp.Infos
	respBody["Page"] = resp.Page
	respBody["StorageKey"] = resp.StorageKey
	respBody["TotalPages"] = resp.TotalPages
	respBody["TotalTokens"] = resp.TotalTokens
	respBody["PageLimit"] = resp.PageLimit
}

// TODO: this method from nft-meta/pkg/imageconvert/utils.go that will be reconstruct
// converte http request with image file to vector
func ImgReqConvertVector(r *http.Request) ([]float32, error) {
	// get file info
	file, handler, err := r.FormFile(UploadFileFeild)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// rewrite file to new request-body
	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)
	fileWriter, err := bodyWriter.CreateFormFile("upload", handler.Filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	bodyWriter.Close()

	icURL := fmt.Sprintf("%v/img2vector/file", config.GetConfig().ImageConverter.Address)
	method := "POST"

	// build request for image-converter
	req, err := http.NewRequestWithContext(r.Context(), method, icURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body1, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// parse response
	resp := &Img2VectorResp{}
	err = json.Unmarshal(body1, resp)
	if err != nil {
		return nil, err
	}

	return resp.Vector, nil
}

type Server struct {
	entranceproto.UnimplementedManagerServer
}

func (s *Server) GetTransfer(ctx context.Context, in *rankerproto.SearchPageRequest) (*rankerproto.SearchResponse, error) {
	token.UseCloudProxyCC()
	return token.SearchPage(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entranceproto.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entranceproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
