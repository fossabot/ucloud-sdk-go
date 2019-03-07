package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
)

const region = "cn-bj2"
const zone = "cn-bj2-05"
const imageID = "uimage-kg0w4u"

var uhostClient *uhost.UHostClient
var unetClient *unet.UNetClient
var ulbClient *ulb.ULBClient

func init() {
	cfg := ucloud.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.Region = region
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	uhostClient = uhost.NewClient(&cfg, &credential)
	unetClient = unet.NewClient(&cfg, &credential)
	ulbClient = ulb.NewClient(&cfg, &credential)

	log.Info("setup clients ...")
}

func main() {
	start := time.Now()
	fmt.Printf("Start: %s\n", start)

	uhostIDs, errs := createUHostBatch(10)
	if len(errs) > 0 {
		log.Error(errs)
	}

	// teardown
	defer deleteUHostBatch(uhostIDs)

	ulbID, err := createULB()
	if err != nil {
		log.Error(err)
	}

	// teardown
	defer deleteULB(ulbID)

	vserverID, err := createVServer(ulbID)
	if err != nil {
		log.Error(err)
	}

	// teardown
	defer deleteVServer(ulbID, vserverID)

	backendIDs, errs := allocateBackendBatch(ulbID, vserverID, uhostIDs)
	if len(errs) > 0 {
		log.Error(errs)
	}

	// teardown
	defer releaseBackendBatch(ulbID, vserverID, backendIDs)

	end := time.Now()
	fmt.Printf("End: %s\n", end)

	fmt.Println(end.Sub(start))
}

func createULB() (string, error) {
	req := ulbClient.NewCreateULBRequest()
	req.Tag = ucloud.String("sdk-example")

	resp, err := ulbClient.CreateULB(req)
	if err != nil {
		return "", err
	}

	time.Sleep(2 * time.Second)

	descReq := ulbClient.NewDescribeULBRequest()
	descReq.Region = ucloud.String(region)
	descReq.Zone = ucloud.String(zone)
	descReq.ULBId = ucloud.String(resp.ULBId)

	_, err = ulbClient.DescribeULB(descReq)
	if err != nil {
		return "", err
	}

	return resp.ULBId, nil
}

func deleteULB(ulbID string) error {
	req := ulbClient.NewDeleteULBRequest()
	req.ULBId = ucloud.String(ulbID)

	_, err := ulbClient.DeleteULB(req)
	if err != nil {
		return err
	}

	return nil
}

func createVServer(id string) (string, error) {
	req := ulbClient.NewCreateVServerRequest()
	req.ULBId = ucloud.String(id)
	// req.Method = ucloud.String("ConsistentHash")

	resp, err := ulbClient.CreateVServer(req)
	if err != nil {
		return "", err
	}

	time.Sleep(5 * time.Second)

	descReq := ulbClient.NewDescribeVServerRequest()
	descReq.Zone = ucloud.String(zone)
	descReq.ULBId = ucloud.String(id)
	descReq.VServerId = ucloud.String(resp.VServerId)

	_, err = ulbClient.DescribeVServer(descReq)
	if err != nil {
		return "", err
	}

	return resp.VServerId, nil
}

func deleteVServer(ulbID, vserverID string) error {
	req := ulbClient.NewDeleteVServerRequest()
	req.ULBId = ucloud.String(ulbID)
	req.VServerId = ucloud.String(vserverID)

	_, err := ulbClient.DeleteVServer(req)
	if err != nil {
		return err
	}

	return nil
}

func allocateBackendBatch(ulbID, vserverID string, uhostIDs []string) (ids []string, errors []error) {
	for _, uhostID := range uhostIDs {
		id, err := allocateBackend(ulbID, vserverID, uhostID)
		if err != nil {
			errors = append(errors, err)
		} else {
			ids = append(ids, id)
		}
	}
	return
}

func allocateBackend(ulbID, vserverID, uhostID string) (string, error) {
	req := ulbClient.NewAllocateBackendRequest()
	req.ULBId = ucloud.String(ulbID)
	req.VServerId = ucloud.String(vserverID)
	req.ResourceType = ucloud.String("UHost")
	req.ResourceId = ucloud.String(uhostID)
	req.Port = ucloud.Int(80)

	resp, err := ulbClient.AllocateBackend(req)
	if err != nil {
		return "", err
	}

	time.Sleep(5 * time.Second)

	descReq := ulbClient.NewDescribeVServerRequest()
	descReq.Zone = ucloud.String(zone)
	descReq.ULBId = ucloud.String(ulbID)
	descReq.VServerId = ucloud.String(vserverID)

	_, err = ulbClient.DescribeVServer(descReq)
	if err != nil {
		return "", err
	}

	return resp.BackendId, nil
}

func releaseBackendBatch(ulbID, vserverID string, backendIDs []string) (errors []error) {
	for _, backendID := range backendIDs {
		err := releaseBackend(ulbID, backendID)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func releaseBackend(ulbID, backendID string) error {
	req := ulbClient.NewReleaseBackendRequest()
	req.ULBId = ucloud.String(ulbID)
	req.BackendId = ucloud.String(backendID)

	_, err := ulbClient.ReleaseBackend(req)
	if err != nil {
		return err
	}

	return nil
}

func createUHostBatch(count int) (ids []string, errors []error) {
	for i := 0; i < count; i++ {
		id, err := createUHost(fmt.Sprintf("sdk-example-%d", i))
		if err != nil {
			errors = append(errors, err)
		} else {
			ids = append(ids, id)
		}
	}

	// wait all uhost instance is running
	if len(ids) > 0 {
		if err := waitForState(ids, uhost.StateRunning); err != nil {
			errors = append(errors, err)
		}
	}

	return
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

	time.Sleep(5 * time.Second)

	descReq := uhostClient.NewDescribeUHostInstanceRequest()
	descReq.Zone = ucloud.String(zone)
	descReq.UHostIds = resp.UHostIds

	_, err = uhostClient.DescribeUHostInstance(descReq)
	if err != nil {
		return "", err
	}

	return resp.UHostIds[0], nil
}

func waitForState(ids []string, state uhost.State) error {
	wait := uhostClient.NewWaitUntilUHostInstanceStateRequest()
	wait.Interval = ucloud.TimeDuration(time.Second * 10)
	wait.MaxAttempts = ucloud.Int(10)
	wait.State = state
	wait.IgnoreError = ucloud.Bool(true)
	desc := uhostClient.NewDescribeUHostInstanceRequest()
	desc.UHostIds = ids
	wait.DescribeRequest = desc

	err := uhostClient.WaitUntilUHostInstanceState(wait)
	if err != nil {
		return err
	}
	return nil
}

func deleteUHostBatch(ids []string) (errors []error) {
	for _, id := range ids {
		err := stopUHost(id)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if err := waitForState(ids, uhost.StateStopped); err != nil {
		errors = append(errors, err)
	}

	for _, id := range ids {
		err := deleteUHost(id)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return
}

func stopUHost(id string) error {
	stop := uhostClient.NewStopUHostInstanceRequest()
	stop.UHostId = ucloud.String(id)
	stop.WithRetry(2)

	_, err := uhostClient.StopUHostInstance(stop)
	if err != nil {
		return err
	}

	return nil
}

func deleteUHost(id string) error {
	req := uhostClient.NewTerminateUHostInstanceRequest()
	req.UHostId = ucloud.String(id)

	_, err := uhostClient.TerminateUHostInstance(req)
	if err != nil {
		return err
	}

	return nil
}
