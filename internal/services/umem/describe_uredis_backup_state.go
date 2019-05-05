//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeURedisBackupState

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeURedisBackupStateRequest is request schema for DescribeURedisBackupState action
type DescribeURedisBackupStateRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 资源id
	GroupId *string `required:"true"`

	// 备份id
	BackupId *string `required:"true"`
}

// DescribeURedisBackupStateResponse is response schema for DescribeURedisBackupState action
type DescribeURedisBackupStateResponse struct {
	response.CommonBase

	// 返回码
	RetCode int

	// 操作名称
	Action string

	// 备份状态
	State string
}

// NewDescribeURedisBackupStateRequest will create request of DescribeURedisBackupState action.
func (c *UMemClient) NewDescribeURedisBackupStateRequest() *DescribeURedisBackupStateRequest {
	req := &DescribeURedisBackupStateRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeURedisBackupState - 查询备份任务的状态
func (c *UMemClient) DescribeURedisBackupState(req *DescribeURedisBackupStateRequest) (*DescribeURedisBackupStateResponse, error) {
	var err error
	var res DescribeURedisBackupStateResponse

	err = c.Client.InvokeAction("DescribeURedisBackupState", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
