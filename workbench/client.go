package workbench

import (
	"fmt"

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
	SetClient(client *openapi.Client)
	GetEndpoint(regionID *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (result *string, err error)
	LoginInstance(request *models.LoginInstanceRequest) (resp *models.LoginInstanceResponseBody, err error)
	LoginInstanceWithOptions(request *models.LoginInstanceRequest, runtime *util.RuntimeOptions) (resp *models.LoginInstanceResponseBody, err error)
}

func NewWorkbenchClient(config *openapi.Config) (workbench WorkbenchClient, err error) {
	client := openapi.Client{}
	if err := client.Init(config); err != nil {
		return nil, err
	}

	workbench = &Workbench{}

	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":  tea.String("ecs-workbench.cn-beijing.aliyuncs.com"),
		"cn-hangzhou": tea.String("ecs-workbench.cn-hangzhou.aliyuncs.com"),
	}

	if err = client.CheckConfig(config); err != nil {
		return nil, err
	}
	if endpoint, err := workbench.GetEndpoint(client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint); err != nil {
		return nil, err
	} else {
		client.Endpoint = endpoint
	}

	workbench.SetClient(&client)

	return workbench, nil
}

func (c *Workbench) GetEndpoint(regionID *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (result *string, err error) {
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

func (c *Workbench) LoginInstance(request *models.LoginInstanceRequest) (resp *models.LoginInstanceResponseBody, err error) {
	runtime := &util.RuntimeOptions{}
	resp, err = c.LoginInstanceWithOptions(request, runtime)
	return
}

func (c *Workbench) LoginInstanceWithOptions(request *models.LoginInstanceRequest, runtime *util.RuntimeOptions) (resp *models.LoginInstanceResponseBody, err error) {
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
		Version:     tea.String("2022-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}

	resp = &models.LoginInstanceResponseBody{}
	fmt.Println(req)
	body, err := c.Client.CallApi(params, req, runtime)
	if err != nil {
		return
	}

	fmt.Println(body)
	err = tea.Convert(body, &resp)

	return
}

func (c *Workbench) SetClient(client *openapi.Client) {
	c.Client = *client
}
