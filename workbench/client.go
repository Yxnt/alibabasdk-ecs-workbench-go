package workbench

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/yxnt/alibabasdk-ecs-workbench-go/workbench/models"
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

func (c *Workbench) getEndpoint(regionID *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (result *string, err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		result = endpoint
		return result, err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionID)])) {
		result = endpointMap[tea.StringValue(regionID)]
		return result, err
	}

	_body, _err := endpointutil.GetEndpointRules(&PRODUCT_ID, regionID, endpointRule, network, suffix)
	if _err != nil {
		return result, err
	}
	result = _body
	return
}

func (c *Workbench) LoginInstance(request *models.LoginInstanceRequest) (resp *models.LoginInstanceResponse, err error) {
	runtime := &util.RuntimeOptions{}
	resp, err = c.LoginInstanceWithOptions(request, runtime)
	return
}

func (c *Workbench) LoginInstanceWithOptions(request *models.LoginInstanceRequest, runtime *util.RuntimeOptions) (resp *models.LoginInstanceResponse, err error) {
	if err = util.ValidateModel(request); err != nil {
		return
	}

	query := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.InstanceLoginInfo)) {
		query["InstanceLoginInfo"] = request.InstanceLoginInfo
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerInfo)) {
		query["PartnerInfo"] = request.PartnerInfo
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccount)) {
		query["UserAccount"] = request.UserAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}

	params := &openapi.Params{
		Action:      tea.String("LoginInstance"),
		Version:     tea.String("2020-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}

	resp = &models.LoginInstanceResponse{}
	body, err := c.Client.CallApi(params, req, runtime)
	if err != nil {
		return
	}

	err = tea.Convert(body, &resp)

	return
}
