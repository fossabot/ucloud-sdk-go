//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ModifyUMemSpaceName

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ModifyUMemSpaceNameRequest is request schema for ModifyUMemSpaceName action
type ModifyUMemSpaceNameRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// UMem内存空间ID
	SpaceId *string `required:"true"`

	// 新的名称,长度(6<=size<=63)
	Name *string `required:"true"`
}

// ModifyUMemSpaceNameResponse is response schema for ModifyUMemSpaceName action
type ModifyUMemSpaceNameResponse struct {
	response.CommonBase
}

// NewModifyUMemSpaceNameRequest will create request of ModifyUMemSpaceName action.
func (c *UMemClient) NewModifyUMemSpaceNameRequest() *ModifyUMemSpaceNameRequest {
	req := &ModifyUMemSpaceNameRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ModifyUMemSpaceName - 修改UMem内存空间名称
func (c *UMemClient) ModifyUMemSpaceName(req *ModifyUMemSpaceNameRequest) (*ModifyUMemSpaceNameResponse, error) {
	var err error
	var res ModifyUMemSpaceNameResponse

	err = c.client.InvokeAction("ModifyUMemSpaceName", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
