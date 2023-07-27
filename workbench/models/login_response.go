package models

type LoginInstanceResponse struct {
	Headers    map[string]string         `json:"headers"`
	StatusCode int64                     `json:"statusCode"`
	Root       LoginInstanceResponseBody `json:"body"`
}

type LoginInstanceResponseBody struct {
	Code      string                        `json:"code"`
	Message   string                        `json:"message"`
	RequestId string                        `json:"requestId"`
	Root      LoginInstanceResponseBodyRoot `json:"Root"`
	Success   bool                          `json:"success"`
}

type LoginInstanceResponseBodyRoot struct {
	DisposableAccount     LoginInstanceResponseBodyRootDisposableAccount       `json:"DisposableAccount"`
	InstanceLoginInfoList []LoginInstanceResponseBodyRootInstanceLoginInfoList `json:"InstanceLoginInfoList"`
	SessionControl        LoginInstanceResponseBodyRootSessionControl          `json:"SessionControl"`
}

type LoginInstanceResponseBodyRootDisposableAccount struct {
	LoginFormActionUrl string `json:"LoginFormActionUrl"`
	LoginUrl           string `json:"LoginUrl"`
}

type LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView struct {
	DefaultViewUrl string `json:"DefaultViewUrl"`
}

type LoginInstanceResponseBodyRootInstanceLoginInfoList struct {
	InstanceId         string                                                              `json:"InstanceId"`
	InstanceLoginToken string                                                              `json:"InstanceLoginToken"`
	InstanceLoginView  LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView `json:"InstanceLoginView"`
	LoginSuccess       bool                                                                `json:"LoginSuccess"`
}

type LoginInstanceResponseBodyRootSessionControl struct {
	BaseUrl string `json:"BaseUrl"`
}
