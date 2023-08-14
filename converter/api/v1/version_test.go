package v1

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/web3eye-io/Web3Eye/config"

	_ "github.com/NpoolPlatform/go-service-framework/pkg/version"
)

func TestVersion(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	resp, err := cli.R().
		Post(fmt.Sprintf("http://localhost:%v/v1/version", config.GetConfig().Converter.HTTPPort))
	if assert.Nil(t, err) {
		fmt.Println(resp)
		assert.Equal(t, 200, resp.StatusCode())
		// we should compare body, but we cannot do here
		// ver, err := version.GetVersion()
		// assert.NotNil(t, err)
		// assert.Equal(t, ver, string(resp.Body()))
	}
}
