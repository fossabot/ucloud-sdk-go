//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem GetUMemSpaceState

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetUMemSpaceStateRequest is request schema for GetUMemSpaceState action
type GetUMemSpaceStateRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 内存空间ID
	SpaceId *string `required:"true"`
}

// GetUMemSpaceStateResponse is response schema for GetUMemSpaceState action
type GetUMemSpaceStateResponse struct {
	response.CommonBase

	// Starting:创建中 Running:运行中 Fail:失败
	State string
}

// NewGetUMemSpaceStateRequest will create request of GetUMemSpaceState action.
func (c *UMemClient) NewGetUMemSpaceStateRequest() *GetUMemSpaceStateRequest {
	req := &GetUMemSpaceStateRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetUMemSpaceState - 获取UMem内存空间列表
func (c *UMemClient) GetUMemSpaceState(req *GetUMemSpaceStateRequest) (*GetUMemSpaceStateResponse, error) {
	var err error
	var res GetUMemSpaceStateResponse

	err = c.client.InvokeAction("GetUMemSpaceState", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
