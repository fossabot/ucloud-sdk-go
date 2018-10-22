//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DescribeUDBLogPackage

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUDBLogPackageRequest is request schema for DescribeUDBLogPackage action
type DescribeUDBLogPackageRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 分页显示的起始偏移，列表操作则指定
	Offset *int `required:"true"`

	// 分页显示的条目数，列表操作则指定
	Limit *int `required:"true"`

	// 需要列出的备份文件类型，每种文件的值如下 2 : BINLOG\_BACKUP 3 : SLOW\_QUERY\_BACKUP 4 : ERRORLOG\_BACKUP
	Type *int `required:"false"`

	// Types作为Type的补充，支持多值传入，可以获取多个类型的日志记录，如：Types.0=2&Types.1=3
	Types []int `required:"false"`

	// DB实例Id，如果指定，则只获取该db的备份信息
	DBId *string `required:"false"`

	// 过滤条件:起始时间(时间戳)
	BeginTime *int `required:"false"`

	// 过滤条件:结束时间(时间戳)
	EndTime *int `required:"false"`
}

// DescribeUDBLogPackageResponse is response schema for DescribeUDBLogPackage action
type DescribeUDBLogPackageResponse struct {
	response.CommonBase

	// 备份信息 参见LogPackageDataSet
	DataSet []LogPackageDataSet

	// 备份总数，如果指定dbid，则是该db备份总数
	TotalCount int
}

// NewDescribeUDBLogPackageRequest will create request of DescribeUDBLogPackage action.
func (c *UDBClient) NewDescribeUDBLogPackageRequest() *DescribeUDBLogPackageRequest {
	req := &DescribeUDBLogPackageRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDBLogPackage - 列表UDB实例binlog或slowlog或errorlog备份信息
func (c *UDBClient) DescribeUDBLogPackage(req *DescribeUDBLogPackageRequest) (*DescribeUDBLogPackageResponse, error) {
	var err error
	var res DescribeUDBLogPackageResponse

	err = c.client.InvokeAction("DescribeUDBLogPackage", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
