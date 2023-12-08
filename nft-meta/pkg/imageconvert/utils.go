package imageconvert

// TODO: will be reconstructed
import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/config"
	tokenhandler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/token"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
)

var (
	ICServer = fmt.Sprintf("%v:%v",
		config.GetConfig().Transform.Domain,
		config.GetConfig().Transform.HTTPPort,
	)
)

type Img2VectorResp struct {
	Vector  []float32 `json:"vector"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}

type VectorInfo struct {
	ID      string    `json:"id"`
	URL     string    `json:"url"`
	Vector  []float32 `json:"vector"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}

// TODO:will be rewrite
func ImgURLConvertVector(imgURL string) ([]float32, error) {
	icURL := fmt.Sprintf("%v/img2vector/url", ICServer)
	method := "POST"

	// build body for request
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("url", imgURL)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), method, icURL, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// parse response
	resp := &Img2VectorResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("image url convert to vector fail,%v", resp.Msg)
	}

	return resp.Vector, nil
}

// TODO:will be rewrite
func HTTPDealVector(ctx context.Context, info *ent.Token) error {
	if info.VectorState != npool.ConvertState_Waiting.String() {
		return errors.New("vector state is not waiting")
	}

	// gain the vector of token
	vector, err := ImgURLConvertVector(info.ImageURL)
	if err != nil || vector == nil {
		info.VectorState = npool.ConvertState_Failed.String()
		vstate := npool.ConvertState(npool.ConvertState_value[info.VectorState])

		h, err := tokenhandler.NewHandler(ctx,
			tokenhandler.WithID(&info.ID, true),
			tokenhandler.WithVectorState(&vstate, true),
			tokenhandler.WithVectorID(&info.VectorID, true),
		)
		if err != nil {
			logger.Sugar().Error(err)
			return err
		}

		_, err = h.UpdateToken(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			return err
		}
		return err
	}

	err = storeToDBAndMilvus(ctx, info, vector)
	if err != nil {
		logger.Sugar().Error(err)
	}
	return err
}

// TODO:will be rewrite
func storeToDBAndMilvus(ctx context.Context, info *ent.Token, vector []float32) (err error) {
	if len(vector) != 0 {
		milvusmgr := milvusdb.NewNFTConllectionMGR()

		err = milvusmgr.Delete(ctx, []int64{info.VectorID})
		if err != nil {
			return err
		}

		// store the vector to milvus
		ids, err := milvusmgr.Create(ctx, [][milvusdb.VectorDim]float32{ToArrayVector(vector)})
		if err != nil {
			return err
		}
		info.VectorID = ids[0]
		info.VectorState = npool.ConvertState_Success.String()
	}

	vstate := npool.ConvertState(npool.ConvertState_value[info.VectorState])

	h, err := tokenhandler.NewHandler(ctx,
		tokenhandler.WithID(&info.ID, true),
		tokenhandler.WithVectorState(&vstate, true),
		tokenhandler.WithVectorID(&info.VectorID, true),
	)
	if err != nil {
		return err
	}
	_, err = h.UpdateToken(ctx)

	return
}

// TODO:will be rewrite
// converte http request with image file to vector
func ImgReqConvertVector(r *http.Request) ([]float32, error) {
	// get file info
	file, handler, err := r.FormFile("upload")
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

	icURL := fmt.Sprintf("%v/img2vector/file", ICServer)
	method := "POST"

	// build request for Transform
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

func ToArrayVector(vector []float32) [milvusdb.VectorDim]float32 {
	// store the vector to milvus
	_vector := [milvusdb.VectorDim]float32{}
	copy(_vector[:], vector)
	return _vector
}
