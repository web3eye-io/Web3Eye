package api

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	_ "github.com/NpoolPlatform/go-service-framework/pkg/version"
)

func TestVersion(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	resp, err := cli.R().
		Post("http://localhost:50230/v2/version")
	if assert.Nil(t, err) {
		fmt.Println(resp)
		assert.Equal(t, 200, resp.StatusCode())
		// we should compare body, but we cannot do here
		// ver, err := version.GetVersion()
		// assert.NotNil(t, err)
		// assert.Equal(t, ver, string(resp.Body()))
	}
}
