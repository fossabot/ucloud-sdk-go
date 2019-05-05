//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet UnBindEIP

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UnBindEIPRequest is request schema for UnBindEIP action
type UnBindEIPRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 弹性IP的资源Id
	EIPId *string `required:"true"`

	// 弹性IP请求解绑的资源类型, 枚举值为: uhost: 云主机; ulb, 负载均衡器 upm: 物理机; hadoophost: 大数据集群;fortresshost：堡垒机；udockhost：容器；udhost：私有专区主机；natgw：NAT网关；udb：udb；vpngw：ipsec vpn；ucdr：云灾备；dbaudit：数据库审计；
	ResourceType *string `required:"true"`

	// 弹性IP请求解绑的资源ID
	ResourceId *string `required:"true"`
}

// UnBindEIPResponse is response schema for UnBindEIP action
type UnBindEIPResponse struct {
	response.CommonBase
}

// NewUnBindEIPRequest will create request of UnBindEIP action.
func (c *UNetClient) NewUnBindEIPRequest() *UnBindEIPRequest {
	req := &UnBindEIPRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UnBindEIP - 将弹性IP从资源上解绑
func (c *UNetClient) UnBindEIP(req *UnBindEIPRequest) (*UnBindEIPResponse, error) {
	var err error
	var res UnBindEIPResponse

	err = c.Client.InvokeAction("UnBindEIP", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
