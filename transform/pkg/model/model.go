package model

import (
	"bytes"
	"context"
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
		pyDir := config.GetConfig().Transform.PyDir
		modelStartFile := fmt.Sprintf("%v/main.py", pyDir)

		cmd := exec.Command("python3", modelStartFile)
		err := cmd.Start()
		if err != nil {
			panic(fmt.Sprintf("failed to start image transform model,err %v", err))
		}
		err = cmd.Wait()
		if err != nil {
			logger.Sugar().Infow("Run Model", "model start file", modelStartFile)
			logger.Sugar().Infow("Run Model", "model start cmd", cmd.String())
			panic(fmt.Sprintf("failed to start image transform model,err %v", err))
		}
	}
	panic(fmt.Sprintf("retry to start image transform model %v times, stop retry", modelRestartRetries))
}

type VectorResp struct {
	Vector []float32 `json:"Vector"`
}

type VectorRep struct {
	ImgPath string `json:"ImgPath"`
}

func ToImageVector(ctx context.Context, imgPath string) ([]float32, error) {
	if _, err := os.Stat(imgPath); err != nil {
		return nil, err
	}
	repBody := VectorRep{ImgPath: imgPath}
	jsonBody, err := json.Marshal(repBody)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://127.0.0.1:8888", body)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
