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
	pbJsonMarshaler jsonpb.Marshaler
)

func init() {
	pbJsonMarshaler = jsonpb.Marshaler{
		EmitDefaults: false,
	}
	mux := servermux.AppServerMux()
	mux.HandleFunc("/search/file", Search)

	pages, err := fs.Sub(resource.ResPages, "pages")
	if err != nil {
		log.Fatalf("failed to load pages: %v", err)
	}
	mux.Handle("/", http.FileServer(http.FS(pages)))
}

type SearchResponse struct {
	Msg         string
	Infos       []byte
	Page        uint32
	StorageKey  string
	TotalPages  uint32
	TotalTokens uint32
	Limit       uint32
}

// nolint
func Search(w http.ResponseWriter, r *http.Request) {
	startT := time.Now()
	respBody := SearchResponse{}
	defer func() {
		_respBody, err := json.Marshal(respBody)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("json marshal response body fail, %v", err)))
			return
		}
		_, err = w.Write(_respBody)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("write response body fail, %v", err)))
		}
	}()

	_limit := r.FormValue(LimitFeild)
	limit, err := strconv.ParseUint(_limit, 10, 32)

	if err != nil {
		respBody.Msg = fmt.Sprintf("failed to parse feild Limit %v, %v", _limit, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// judge weather filesize exceed max-size
	err = r.ParseMultipartForm(MaxUploadFileSize)
	if err != nil {
		respBody.Msg = fmt.Sprintf("read file failed %v, %v", MaxUploadFileSize, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inT := time.Now()
	logger.Sugar().Infof("check params %v ms", inT.UnixMilli()-startT.UnixMilli())

	// convert to vector
	vector, err := ImgReqConvertVector(r)
	if err != nil {
		respBody.Msg = fmt.Sprintf("image convert fail, %v", err)
		w.WriteHeader(http.StatusBadRequest)
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
		respBody.Msg = fmt.Sprintf("search fail, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	inT = time.Now()
	logger.Sugar().Infof("finish query id %v ms", inT.UnixMilli()-startT.UnixMilli())

	buff := bytes.NewBuffer([]byte{})
	err = pbJsonMarshaler.Marshal(buff, resp)

	respBody.Msg = fmt.Sprintf("have %v infos", len(resp.Infos))
	respBody.Infos = buff.Bytes()
	respBody.Page = resp.Page
	respBody.StorageKey = resp.StorageKey
	respBody.TotalPages = resp.TotalPages
	respBody.TotalTokens = resp.TotalTokens
	respBody.Limit = resp.Limit
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

	res, err := http.Post(icURL, bodyWriter.FormDataContentType(), body)
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

func (s *Server) SearchPage(ctx context.Context, in *rankerproto.SearchPageRequest) (*rankerproto.SearchResponse, error) {
	token.UseCloudProxyCC()
	return token.SearchPage(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entranceproto.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entranceproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
