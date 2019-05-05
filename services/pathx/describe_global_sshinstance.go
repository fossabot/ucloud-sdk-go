//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api PathX DescribeGlobalSSHInstance

package pathx

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeGlobalSSHInstanceRequest is request schema for DescribeGlobalSSHInstance action
type DescribeGlobalSSHInstanceRequest struct {
	request.CommonBase

	// 实例ID，资源唯一标识
	InstanceId *string `required:"false"`
}

// DescribeGlobalSSHInstanceResponse is response schema for DescribeGlobalSSHInstance action
type DescribeGlobalSSHInstanceResponse struct {
	response.CommonBase

	// GlobalSSH实例列表，实例的属性参考GlobalSSHInfo模型
	InstanceSet []GlobalSSHInfo
}

// NewDescribeGlobalSSHInstanceRequest will create request of DescribeGlobalSSHInstance action.
func (c *PathXClient) NewDescribeGlobalSSHInstanceRequest() *DescribeGlobalSSHInstanceRequest {
	req := &DescribeGlobalSSHInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeGlobalSSHInstance - 获取GlobalSSH实例列表（传实例ID获取单个实例信息，不传获取项目下全部实例）
func (c *PathXClient) DescribeGlobalSSHInstance(req *DescribeGlobalSSHInstanceRequest) (*DescribeGlobalSSHInstanceResponse, error) {
	var err error
	var res DescribeGlobalSSHInstanceResponse

	err = c.Client.InvokeAction("DescribeGlobalSSHInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
