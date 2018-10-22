//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeUMemPrice

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUMemPriceRequest is request schema for DescribeUMemPrice action
type DescribeUMemPriceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 购买umem大小,单位:GB,范围[1~1024]
	Size *int `required:"true"`

	// 空间类型:single(无热备),double(热备)(默认: double)
	Type *string `required:"true"`

	// Year， Month， Dynamic，Trial 如果不指定，则一次性获取三种计费
	ChargeType *string `required:"false"`

	// 购买UMem的时长，默认值为1
	Quantity *int `required:"false"`
}

// DescribeUMemPriceResponse is response schema for DescribeUMemPrice action
type DescribeUMemPriceResponse struct {
	response.CommonBase

	// 价格 参数见 UMemPriceSet
	DataSet []UMemPriceSet
}

// NewDescribeUMemPriceRequest will create request of DescribeUMemPrice action.
func (c *UMemClient) NewDescribeUMemPriceRequest() *DescribeUMemPriceRequest {
	req := &DescribeUMemPriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUMemPrice - 获取UMem实例价格信息
func (c *UMemClient) DescribeUMemPrice(req *DescribeUMemPriceRequest) (*DescribeUMemPriceResponse, error) {
	var err error
	var res DescribeUMemPriceResponse

	err = c.client.InvokeAction("DescribeUMemPrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
