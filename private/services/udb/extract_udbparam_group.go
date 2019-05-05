//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB ExtractUDBParamGroup

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ExtractUDBParamGroupRequest is request schema for ExtractUDBParamGroup action
type ExtractUDBParamGroupRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 配置id
	GroupId *int `required:"true"`

	// 是否跨可用区，RegionFlag为true时表示跨可用区配置文件
	RegionFlag *bool `required:"false"`
}

// ExtractUDBParamGroupResponse is response schema for ExtractUDBParamGroup action
type ExtractUDBParamGroupResponse struct {
	response.CommonBase

	// 配置文件内容
	Content string
}

// NewExtractUDBParamGroupRequest will create request of ExtractUDBParamGroup action.
func (c *UDBClient) NewExtractUDBParamGroupRequest() *ExtractUDBParamGroupRequest {
	req := &ExtractUDBParamGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ExtractUDBParamGroup - 获取配置文件内容
func (c *UDBClient) ExtractUDBParamGroup(req *ExtractUDBParamGroupRequest) (*ExtractUDBParamGroupResponse, error) {
	var err error
	var res ExtractUDBParamGroupResponse

	err = c.Client.InvokeAction("ExtractUDBParamGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
