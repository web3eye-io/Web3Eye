package token

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	converter "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/converter/v1/token"
	crud "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/crud/v1/token"
	res "github.com/web3eye-io/cyber-tracer/nft-meta/resource"
	val "github.com/web3eye-io/cyber-tracer/proto/cybertracer"

	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/milvusdb"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/servermux"
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/token"
)

// 8mb
const MaxUploadFileSize = 1 << 10 << 10 << 3

type Img2VectorResp struct {
	Vector  []float32 `json:"vector"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}

type SearchToken struct {
	npool.Token
	Distance float32
}

func init() {
	mux := servermux.AppServerMux()
	mux.HandleFunc("/search/file", Search)

	pages, err := fs.Sub(res.ResPages, "pages")
	if err != nil {
		log.Fatalf("failed to load pages: %v", err)
	}
	mux.Handle("/", http.FileServer(http.FS(pages)))
}

func Search(w http.ResponseWriter, r *http.Request) {
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

	ctx := context.Background()

	// judge weather filesize exceed max-size
	err := r.ParseMultipartForm(MaxUploadFileSize)
	if err != nil {
		respBody["msg"] = fmt.Sprintf("filesize is much than %v, %v", MaxUploadFileSize, err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// convert to vector
	vector, err := imageconvert.ImgReqConvertVector(r)
	if err != nil {
		respBody["msg"] = fmt.Sprintf("image convert fail, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	milvusmgr := milvusdb.NewNFTConllectionMGR()
	_vector := imageconvert.ToArrayVector(vector)
	// TODO: It should be given by request
	defaultTopN := 10
	_scores, err := milvusmgr.Search(
		context.Background(),
		[][milvusdb.VectorDim]float32{_vector},
		defaultTopN,
	)
	if err != nil {
		respBody["msg"] = fmt.Sprintf("search image vector fail, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(_scores) == 0 {
		respBody["msg"] = "have no vector"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	scores := _scores[0]
	vIDs := []int64{}
	for i := range scores {
		vIDs = append(vIDs, i)
	}

	conds := &npool.Conds{
		VectorIDs: &val.Int64SliceVal{
			Op:    "in",
			Value: vIDs,
		},
	}

	rows, _, err := crud.Rows(ctx, conds, 0, len(vIDs))
	if err != nil {
		respBody["msg"] = fmt.Sprintf("query nft-info fail, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	infos := []*SearchToken{}
	for i := 0; i < len(rows); i++ {
		info := &SearchToken{}
		info.Token = *converter.Ent2Grpc(rows[i])
		info.Distance = scores[info.VectorID]
		infos = append(infos, info)
	}
	respBody["msg"] = fmt.Sprintf("have %v infos", len(infos))
	respBody["data"] = infos
}
