//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeUMemUpgradePrice

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUMemUpgradePriceRequest is request schema for DescribeUMemUpgradePrice action
type DescribeUMemUpgradePriceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 购买UMem大小,单位:GB
	Size *int `required:"true"`

	// 空间类型:single(无热备),double(热备)(默认: double)
	Type *string `required:"true"`

	// 需要升级的空间的SpaceId
	SpaceId *string `required:"true"`
}

// DescribeUMemUpgradePriceResponse is response schema for DescribeUMemUpgradePrice action
type DescribeUMemUpgradePriceResponse struct {
	response.CommonBase

	// 价格
	Price float64
}

// NewDescribeUMemUpgradePriceRequest will create request of DescribeUMemUpgradePrice action.
func (c *UMemClient) NewDescribeUMemUpgradePriceRequest() *DescribeUMemUpgradePriceRequest {
	req := &DescribeUMemUpgradePriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUMemUpgradePrice - 获取UMem升级价格信息
func (c *UMemClient) DescribeUMemUpgradePrice(req *DescribeUMemUpgradePriceRequest) (*DescribeUMemUpgradePriceResponse, error) {
	var err error
	var res DescribeUMemUpgradePriceResponse

	err = c.client.InvokeAction("DescribeUMemUpgradePrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
