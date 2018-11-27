//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ModifyUMemcacheGroupName

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ModifyUMemcacheGroupNameRequest is request schema for ModifyUMemcacheGroupName action
type ModifyUMemcacheGroupNameRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 组的ID
	GroupId *string `required:"true"`

	// 组名称
	Name *string `required:"true"`
}

// ModifyUMemcacheGroupNameResponse is response schema for ModifyUMemcacheGroupName action
type ModifyUMemcacheGroupNameResponse struct {
	response.CommonBase
}

// NewModifyUMemcacheGroupNameRequest will create request of ModifyUMemcacheGroupName action.
func (c *UMemClient) NewModifyUMemcacheGroupNameRequest() *ModifyUMemcacheGroupNameRequest {
	req := &ModifyUMemcacheGroupNameRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ModifyUMemcacheGroupName - 修改UMemcache名称
func (c *UMemClient) ModifyUMemcacheGroupName(req *ModifyUMemcacheGroupNameRequest) (*ModifyUMemcacheGroupNameResponse, error) {
	var err error
	var res ModifyUMemcacheGroupNameResponse

	err = c.client.InvokeAction("ModifyUMemcacheGroupName", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
