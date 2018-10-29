package ucloud

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

var client *Client

func TestMain(m *testing.M) {
	testSetup()
	exitCode := m.Run()
	testTeardown()
	os.Exit(exitCode)
}

func testSetup() {
	cfg := NewConfig()
	// cfg.BaseUrl = "https://api-mock.pre.ucloudadmin.com/?_user=yufei.li%40ucloud.cn"
	cfg.BaseUrl = "https://api.ucloud.cn"
	cfg.Region = "cn-bj2"
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")
	cfg.LogLevel = log.DebugLevel

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	log.Infof("config: %#v, credential: %#v", cfg, credential)
	client = NewClient(&cfg, &credential)
}

func testTeardown() {}

type MockRequest struct {
	request.CommonBase
}

type MockResponse struct {
	response.CommonBase

	TotalCount int
	UHostSet   []map[string]interface{}
}

func TestCommonInvokeAction(t *testing.T) {
	req := &MockRequest{}
	resp := &MockResponse{}

	err := client.InvokeAction("DescribeUHostInstance", client.SetupRequest(req), resp)
	assert.Nil(t, err)
	assert.Condition(t, func() bool { return resp.TotalCount >= 0 })
}

func TestCommonInvokeActionNotFound(t *testing.T) {
	req := &MockRequest{}
	resp := &MockResponse{}
	err := client.InvokeAction("TestApi", client.SetupRequest(req), resp)
	assert.NotNil(t, err)
	uErr, ok := err.(uerr.Error)
	assert.True(t, ok)
	assert.Equal(t, 161, uErr.Code())
	assert.Contains(t, uErr.Message(), "not found")
}

func Test_errorHandler(t *testing.T) {
	req := &MockRequest{}
	resp := &MockResponse{}

	steps := []struct {
		name string
		step func() error
	}{
		{
			name: "unexcepted error",
			step: func() error {
				_, err := errorHandler(client, req, resp, errors.New("unexcepted error"))
				if uErr, ok := err.(uerr.ClientError); !ok || uErr.Name() != uerr.ErrSendRequest {
					return errors.New("unexcepted error should be convert to unknown client error")
				}
				return nil
			},
		},
		{
			name: "http status error",
			step: func() error {
				_, err := errorHandler(client, req, resp, uerr.NewServerStatusError(404, "404 NotFound"))
				if uErr, ok := err.(uerr.ServerError); !ok || uErr.StatusCode() != 404 {
					return errors.New("http status error should be convert to status server error")
				}
				return nil
			},
		},
		{
			name: "timeout error",
			step: func() error {
				httpClient := &http.Client{Timeout: time.Duration(1)}
				httpReq, err := http.NewRequest("GET", "https://httpbin.org/delay/2", nil)
				if err != nil {
					return err
				}
				_, err = httpClient.Do(httpReq)
				_, err = errorHandler(client, req, resp, err)
				if uErr, ok := err.(uerr.ClientError); !ok || uErr.Name() != uerr.ErrNetwork {
					return errors.New("timeout error should be convert to network client error")
				}
				return nil
			},
		},
		{
			name: "bussiness error",
			step: func() error {
				resp := &response.CommonBase{Message: "Missing Action", RetCode: 160}
				_, err := errorHandler(client, req, resp, nil)
				if uErr, ok := err.(uerr.ServerError); !ok || uErr.Code() != 160 {
					return errors.New("ucloud error should be raised for non-zero retCode")
				}
				return nil
			},
		},
	}
	for _, step := range steps {
		t.Run(step.name, func(t *testing.T) {
			err := step.step()
			if err != nil {
				t.Errorf("errorHandler() error %s", err)
			}
		})
	}
}
