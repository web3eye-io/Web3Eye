package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/config"
)

var modelRestartRetries = 3

func Run() {
	for i := 0; i < modelRestartRetries; i++ {
		pyDir := config.GetConfig().Converter.PyDir
		modelStartFile := fmt.Sprintf("%vmodel/main.py", pyDir)
		fmt.Println(modelStartFile)

		cmd := exec.Command("python3", modelStartFile)
		err := cmd.Start()
		if err != nil {
			logger.Sugar().Errorf("failed to start image converter model")
		}
		cmd.Wait()
	}
	panic(fmt.Sprintf("retry to start image converter model %v times, stop retry", modelRestartRetries))
}

type VectorResp struct {
	Vector []float32 `json:"Vector"`
}

type VectorRep struct {
	ImgPath string `json:"ImgPath"`
}

func ToImageVector(imgPath string) ([]float32, error) {
	if _, err := os.Stat(imgPath); err != nil {
		return nil, err
	}
	repBody := VectorRep{ImgPath: imgPath}
	jsonBody, err := json.Marshal(repBody)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonBody)
	resp, err := http.Post("http://127.0.0.1:8888", "application/json", body)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	VectorResp := VectorResp{}

	err = json.Unmarshal(respBody, &VectorResp)
	if err != nil {
		return nil, err
	}

	return VectorResp.Vector, nil
}
