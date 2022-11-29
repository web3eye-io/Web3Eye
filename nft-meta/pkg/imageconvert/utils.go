package imageconvert

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/cyber-tracer/message/cybertracer/nftmeta/v1/token"

	"github.com/google/uuid"
	crud "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/milvusdb"
)

var (
	ICServer = "http://172.16.31.31:8080"
)

type Img2VectorResp struct {
	Vector  []float32 `json:"vector"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}

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

func DealVectorState(ctx context.Context, id uuid.UUID) {
	// query record and check it`s vector_state
	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Warn(err)
		return
	}

	if info.VectorState != npool.ConvertState_Waiting.String() {
		return
	}
	infoID := info.ID.String()

	// gain the vector of token
	vector, err := ImgURLConvertVector(info.ImageURL)
	if err != nil || vector == nil {
		info.VectorState = npool.ConvertState_Failed.String()
		vstate := npool.ConvertState(npool.ConvertState_value[info.VectorState])
		_, err := crud.Update(ctx, &npool.TokenReq{
			ID:          &infoID,
			VectorState: &vstate,
			VectorID:    &info.VectorID,
		})
		if err != nil {
			logger.Sugar().Warn(err)
			return
		}
		return
	}

	milvusmgr := milvusdb.NewNFTConllectionMGR()
	err = milvusmgr.Delete(ctx, []int64{info.VectorID})
	if err != nil {
		logger.Sugar().Warn(err)
		return
	}

	// store the vector to milvus
	ids, err := milvusmgr.Create(ctx, [][milvusdb.VectorDim]float32{ToArrayVector(vector)})
	if err != nil {
		logger.Sugar().Warn(err)
		return
	}

	// update token record to database
	info.VectorState = npool.ConvertState_Success.String()
	info.VectorID = ids[0]
	vstate := npool.ConvertState(npool.ConvertState_value[info.VectorState])
	_, err = crud.Update(ctx, &npool.TokenReq{
		ID:          &infoID,
		VectorState: &vstate,
		VectorID:    &info.VectorID,
	})
	if err != nil {
		logger.Sugar().Warn(err)
		return
	}
}

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

func ToArrayVector(vector []float32) [milvusdb.VectorDim]float32 {
	// store the vector to milvus
	_vector := [milvusdb.VectorDim]float32{}
	copy(_vector[:], vector)
	return _vector
}
