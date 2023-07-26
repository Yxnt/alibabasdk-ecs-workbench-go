package test

import (
	"testing"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/yxnt/alibabasdk-ecs-workbench-go/workbench"
)

func Test_GetEndpoint(t *testing.T) {
	regionID := "cn-shanghai"

	config := &openapi.Config{
		RegionId: &regionID,
	}
	if workbenchClient, err := workbench.NewWorkbenchClient(config); err != nil {
		t.Errorf("Init workbench client failed: %v", err)
	} else {
		if workbenchClient.Client.RegionId != config.RegionId {
			t.Errorf("region id not match: %v", workbenchClient.Client.RegionId)
		}

		if workbenchClient.Client.Endpoint == nil {
			t.Errorf("endpoint is nil")
		} else {
			t.Log(*workbenchClient.Client.Endpoint)
		}
	}
}
