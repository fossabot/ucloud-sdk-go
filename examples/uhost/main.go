package main

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"os"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
)

const region = "cn-bj2"
const zone = "cn-bj2-05"
const imageID = "uimage-kg0w4u"

var uhostClient *uhost.UHostClient
var unetClient *unet.UNetClient

func main() {
	cfg, credential := loadConfig()
	uhostClient = uhost.NewClient(cfg, credential)
	unetClient = unet.NewClient(cfg, credential)

	eipID, err := createEIP("sdk-example-uhost")
	if err != nil {
		panic(err)
	}

	uhostID, err := createUHost("sdk-example-uhost")
	if err != nil {
		panic(err)
	}

	err = bindEIPToUHost(eipID, uhostID)
	if err != nil {
		panic(err)
	}
}

func loadConfig() (*ucloud.Config, *auth.Credential) {
	cfg := ucloud.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.Region = region
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	log.Info("setup clients ...")

	return &cfg, &credential
}

func bindEIPToUHost(eipID, uhostID string) error {
	req := unetClient.NewBindEIPRequest()
	req.EIPId = ucloud.String(eipID)
	req.ResourceId = ucloud.String(uhostID)
	req.ResourceType = ucloud.String("uhost")

	_, err := unetClient.BindEIP(req)
	if err != nil {
		return err
	}
	return nil
}

func createEIP(name string) (string, error) {
	req := unetClient.NewAllocateEIPRequest()
	req.Name = ucloud.String(name)
	req.Bandwidth = ucloud.Int(2)
	req.OperatorName = ucloud.String("Bgp")

	resp, err := unetClient.AllocateEIP(req)
	if err != nil {
		return "", err
	}

	return resp.EIPSet[0].EIPId, nil
}

func createUHost(name string) (string, error) {
	req := uhostClient.NewCreateUHostInstanceRequest()
	req.Name = ucloud.String(name)
	req.Zone = ucloud.String(zone)       // TODO: use random zone
	req.ImageId = ucloud.String(imageID) // TODO: use random image
	req.LoginMode = ucloud.String("Password")
	req.Password = ucloud.String("somePassword_")
	req.ChargeType = ucloud.String("Dynamic")
	req.CPU = ucloud.Int(1)
	req.Memory = ucloud.Int(1024)
	req.Tag = ucloud.String("sdk-example")

	resp, err := uhostClient.CreateUHostInstance(req)
	if err != nil {
		return "", err
	}

	return resp.UHostIds[0], nil
}
