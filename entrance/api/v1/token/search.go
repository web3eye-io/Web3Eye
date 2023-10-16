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
	"github.com/gogo/protobuf/jsonpb"

	"github.com/web3eye-io/Web3Eye/common/servermux"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/entrance/resource"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	"github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/token"

	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

// 8mb

const (
	MaxUploadFileSize = 1 << 10 << 10 << 3
	UploadFileFeild   = "UploadFile"
	LimitFeild        = "Limit"
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

var (
	pbJSONMarshaler jsonpb.Marshaler
)

func init() {
	mux := servermux.AppServerMux()
	mux.HandleFunc("/search/file", Search)

	pages, err := fs.Sub(resource.ResPages, "pages")
	if err != nil {
		log.Fatalf("failed to load pages: %v", err)
	}
	mux.Handle("/", http.FileServer(http.FS(pages)))
}

func Search(w http.ResponseWriter, r *http.Request) {
	startT := time.Now()
	respBody := []byte{}
	var err error
	var errMsg string
	defer func() {
		if errMsg != "" {
			logger.Sugar().Infof("failed to search, err: %v", errMsg)
			w.WriteHeader(http.StatusBadRequest)
			respBody = []byte(errMsg)
		}

		_, err = w.Write(respBody)
		if err != nil {
			logger.Sugar().Errorf("failed to write response,err %v", err)
		}
	}()

	_limit := r.FormValue(LimitFeild)
	baseNum := 10
	bitSize := 32
	limit, err := strconv.ParseUint(_limit, baseNum, bitSize)

	if err != nil {
		errMsg = fmt.Sprintf("failed to parse feild Limit %v, %v", _limit, err)
		return
	}

	// judge weather filesize exceed max-size
	err = r.ParseMultipartForm(MaxUploadFileSize)
	if err != nil {
		errMsg = fmt.Sprintf("read file failed %v, %v", MaxUploadFileSize, err)
		return
	}

	inT := time.Now()
	logger.Sugar().Infof("check params %v ms", inT.UnixMilli()-startT.UnixMilli())

	// convert to vector
	vector, err := ImgReqConvertVector(r.Context(), r)
	if err != nil {
		errMsg = fmt.Sprintf("image convert fail, %v", err)
		return
	}

	inT = time.Now()
	logger.Sugar().Infof("finish convert to vector %v ms", inT.UnixMilli()-startT.UnixMilli())

	token.UseCloudProxyCC()
	resp, err := token.Search(context.Background(), &rankerproto.SearchTokenRequest{
		Vector: vector,
		Limit:  uint32(limit),
	})
	if err != nil {
		errMsg = fmt.Sprintf("search fail, %v", err)
		return
	}

	inT = time.Now()
	logger.Sugar().Infof("finish query id %v ms", inT.UnixMilli()-startT.UnixMilli())

	pbJSONMarshaler.EmitDefaults = true
	buff := bytes.NewBuffer([]byte{})
	err = pbJSONMarshaler.Marshal(buff, resp)
	if err != nil {
		errMsg = fmt.Sprintf("marshal result fail, %v", err)
		return
	}

	respBody = buff.Bytes()
}

// converte http request with image file to vector
func ImgReqConvertVector(ctx context.Context, r *http.Request) ([]float32, error) {
	// get file info
	file, handler, err := r.FormFile(UploadFileFeild)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// rewrite file to new request-body
	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)
	fileWriter, err := bodyWriter.CreateFormFile(UploadFileFeild, handler.Filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	bodyWriter.Close()
	ICServer := fmt.Sprintf("%v:%v",
		config.GetConfig().Transform.Domain,
		config.GetConfig().Transform.HTTPPort,
	)
	icURL := fmt.Sprintf("http://%v/v1/transform/file", ICServer)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, icURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body1, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse response
	vectorResp := &Img2VectorResp{}
	err = json.Unmarshal(body1, vectorResp)
	if err != nil {
		return nil, err
	}

	return vectorResp.Vector, nil
}

func (s *Server) SearchPage(ctx context.Context, in *rankerproto.SearchPageRequest) (*rankerproto.SearchResponse, error) {
	token.UseCloudProxyCC()
	return token.SearchPage(ctx, in)
}
