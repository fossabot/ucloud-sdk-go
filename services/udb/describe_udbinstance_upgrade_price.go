//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DescribeUDBInstanceUpgradePrice

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUDBInstanceUpgradePriceRequest is request schema for DescribeUDBInstanceUpgradePrice action
type DescribeUDBInstanceUpgradePriceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 实例的Id
	DBId *string `required:"true"`

	// 内存限制(MB)
	MemoryLimit *int `required:"true"`

	// 磁盘空间(GB), 暂时支持20G - 500G
	DiskSpace *int `required:"true"`

	// 是否使用SSD，默认为false
	UseSSD *bool `required:"false"`

	// SSD类型，可选值为"SATA"、"PCI-E"，如果UseSSD为true ，则必选
	SSDType *string `required:"false"`
}

// DescribeUDBInstanceUpgradePriceResponse is response schema for DescribeUDBInstanceUpgradePrice action
type DescribeUDBInstanceUpgradePriceResponse struct {
	response.CommonBase

	// 价格，单位分
	Price float64
}

// NewDescribeUDBInstanceUpgradePriceRequest will create request of DescribeUDBInstanceUpgradePrice action.
func (c *UDBClient) NewDescribeUDBInstanceUpgradePriceRequest() *DescribeUDBInstanceUpgradePriceRequest {
	req := &DescribeUDBInstanceUpgradePriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDBInstanceUpgradePrice - 获取UDB实例升降级价格信息
func (c *UDBClient) DescribeUDBInstanceUpgradePrice(req *DescribeUDBInstanceUpgradePriceRequest) (*DescribeUDBInstanceUpgradePriceResponse, error) {
	var err error
	var res DescribeUDBInstanceUpgradePriceResponse

	err = c.client.InvokeAction("DescribeUDBInstanceUpgradePrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
