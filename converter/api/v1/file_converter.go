package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"

	"github.com/web3eye-io/Web3Eye/common/servermux"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/converter/pkg/model"
)

const (
	UploadFileFeild = "UploadFile"
)

type ConverterResp struct {
	Vector []float32
	Msg    string
}

func init() {
	mux := servermux.AppServerMux()
	mux.HandleFunc("/converter/file", ConverterFile)
}

func ConverterFile(w http.ResponseWriter, r *http.Request) {
	startT := time.Now()
	resp := ConverterResp{}
	defer func() {
		_respBody, err := json.Marshal(resp)
		if err != nil {
			logger.Sugar().Errorf("json marshal response body fail, %v", err)
			_respBody = []byte("{'Vector':null,'Msg':'wrong server state,connot marshal response'}")
		}

		_, err = w.Write(_respBody)
		if err != nil {
			logger.Sugar().Errorf("write response body fail, %v", err)
		}
	}()

	inT := time.Now()
	logger.Sugar().Infof("check params %v ms", inT.UnixMilli()-startT.UnixMilli())

	// get file info
	file, handler, err := r.FormFile(UploadFileFeild)
	if err != nil {
		errStr := err.Error()
		logger.Sugar().Errorf("failed to get file info ,err: %v", err)
		resp.Msg = errStr
		return
	}
	defer file.Close()

	ext := path.Ext(handler.Filename)

	fmt.Println(ext)
	// write to file
	filePath := fmt.Sprintf("%v/%v%v", config.GetConfig().Converter.DataDir, uuid.NewString(), ext)
	fileContent, err := os.Create(filePath)
	if err != nil {
		errStr := err.Error()
		logger.Sugar().Errorf("failed to create file ,err: %v", err)
		resp.Msg = errStr
		return
	}

	_, err = io.Copy(fileContent, file)
	if err != nil {
		errStr := err.Error()
		logger.Sugar().Errorf("failed to create file ,err: %v", err)
		resp.Msg = errStr
		return
	}
	// defer os.Remove(filePath)

	vector, err := model.ToImageVector(filePath)
	if err != nil {
		errStr := err.Error()
		logger.Sugar().Errorf("failed to get file info ,err: %v", err)
		resp.Msg = errStr
		return
	}

	if vector != nil {
		resp.Vector = vector
	}

	inT = time.Now()
	logger.Sugar().Infof("finish convert to vector %v ms", inT.UnixMilli()-startT.UnixMilli())
}
