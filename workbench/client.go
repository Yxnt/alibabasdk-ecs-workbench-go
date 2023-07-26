package workbench

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

var (
	PRODUCT_ID string = "ecs-workbench"
)

type Workbench struct {
	Client openapi.Client
}

type WorkbenchClient interface {
	LoginInstance()
	LoginInstanceWithOptions()
}

func NewWorkbenchClient(config *openapi.Config) (workbench *Workbench, err error) {
	client := openapi.Client{}
	workbench = &Workbench{}

	if endpoint, err := workbench.getEndpoint(client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint); err != nil {
		return nil, err
	} else {
		config.Endpoint = endpoint
	}

	if err = client.CheckConfig(config); err != nil {
		return nil, err
	}

	if err := client.Init(config); err != nil {
		return nil, err
	}

	workbench.Client = client

	return
}

func (c *Workbench) getEndpoint(regionID *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionID)])) {
		_result = endpointMap[tea.StringValue(regionID)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(&PRODUCT_ID, regionID, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (c *Workbench) LoginInstance() {
	return
}

func (c *Workbench) LoginInstanceWithOptions() {
	return
}
