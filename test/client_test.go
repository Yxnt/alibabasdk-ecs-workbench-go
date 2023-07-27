package test

import (
	"testing"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/yxnt/alibabasdk-ecs-workbench-go/workbench"
	"github.com/yxnt/alibabasdk-ecs-workbench-go/workbench/models"
)

func Test_GetEndpoint(t *testing.T) {
	config := &openapi.Config{
		RegionId: tea.String("cn-hangzhou"),
	}
	if workbenchClient, err := workbench.NewWorkbenchClient(config); err != nil {
		t.Errorf("Init workbench client failed: %v", err)
	} else {
		if endpoint, err := workbenchClient.GetEndpoint(config.RegionId, nil, config.Network, config.Suffix, nil, config.Endpoint); err != nil {
			t.Errorf("Get endpoint failed: %v", err)
		} else {
			t.Logf("Endpoint: %v", *endpoint)
		}
	}
}

func Test_ApiCall(t *testing.T) {
	config := &openapi.Config{
		RegionId:        tea.String("cn-hangzhou"),
		AccessKeyId:     tea.String(""),
		AccessKeySecret: tea.String(""),
	}
	workbenchClient, err := workbench.NewWorkbenchClient(config)
	if err != nil {
		t.Errorf("Init workbench client failed: %v", err)
	}

	req := models.LoginInstanceRequest{
		InstanceLoginInfo: models.LoginInstanceRequestInstanceLoginInfo{
			InstanceId: "xxxxx",
			Protocol:   "ssh",
			Username:   "root",
			Password:   "123456",
		},
	}
	resp, err := workbenchClient.LoginInstance(&req)
	if err != nil {
		t.Error(err)
	}
	if !resp.Success {
		t.Errorf("Login instance failed: %s", resp.Message)
	}
	t.Log(resp.Root.InstanceLoginInfoList[0].InstanceLoginView.DefaultViewUrl)
}
