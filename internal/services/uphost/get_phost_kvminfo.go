//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UPHost GetPHostKVMInfo

package uphost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetPHostKVMInfoRequest is request schema for GetPHostKVMInfo action
type GetPHostKVMInfoRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// GetPHostKVMInfoResponse is response schema for GetPHostKVMInfo action
type GetPHostKVMInfoResponse struct {
	response.CommonBase

	// 资源id
	PHostId string

	// kvm配置文件
	PHostKVMConf string

	// 配置文件类型
	PHostConfType string

	// 代理ip
	ProxyIP string
}

// NewGetPHostKVMInfoRequest will create request of GetPHostKVMInfo action.
func (c *UPHostClient) NewGetPHostKVMInfoRequest() *GetPHostKVMInfoRequest {
	req := &GetPHostKVMInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetPHostKVMInfo - 控制台紧急登录功能
func (c *UPHostClient) GetPHostKVMInfo(req *GetPHostKVMInfoRequest) (*GetPHostKVMInfoResponse, error) {
	var err error
	var res GetPHostKVMInfoResponse

	err = c.Client.InvokeAction("GetPHostKVMInfo", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
