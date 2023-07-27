package models

type LoginInstanceRequest struct {
	InstanceLoginInfo LoginInstanceRequestInstanceLoginInfo `json:"InstanceLoginInfo,omitempty"`
	PartnerInfo       LoginInstanceRequestPartnerInfo       `json:"PartnerInfo,omitempty"`
	UserAccount       LoginInstanceRequestUserAccount       `json:"UserAccount,omitempty"`
}

type LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo struct {
	ClusterId     string `json:"ClusterId,omitempty"`
	ClusterName   string `json:"ClusterName,omitempty"`
	ContainerName string `json:"ContainerName,omitempty"`
	Deployment    string `json:"Deployment,omitempty"`
	Endpoint      string `json:"Endpoint,omitempty"`
	Headers       string `json:"Headers,omitempty"`
	Namespace     string `json:"Namespace,omitempty"`
	PodName       string `json:"PodName,omitempty"`
}

type LoginInstanceRequestInstanceLoginInfoOptions struct {
	AudioMuteSeconds                 int64                                                     `json:"AudioMuteSeconds,omitempty"`
	ContainerInfo                    LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo `json:"ContainerInfo,omitempty"`
	NotificationEventTypes           string                                                    `json:"NotificationEventTypes,omitempty"`
	NotificationRecipientUrl         string                                                    `json:"NotificationRecipientUrl,omitempty"`
	NotificationRetryIntervalSeconds string                                                    `json:"NotificationRetryIntervalSeconds,omitempty"`
	NotificationRetryLimit           string                                                    `json:"NotificationRetryLimit,omitempty"`
	OperationDisableSeconds          int64                                                     `json:"OperationDisableSeconds,omitempty"`
	SessionControl                   string                                                    `json:"SessionControl,omitempty"`
	VideoFreezeSeconds               string                                                    `json:"VideoFreezeSeconds,omitempty"`
}

type LoginInstanceRequestInstanceLoginInfo struct {
	AuthenticationType string `json:"AuthenticationType,omitempty"`
	Certificate        string `json:"Certificate,omitempty"`
	DurationSeconds    string `json:"DurationSeconds,omitempty"`
	ExpireTime         string `json:"ExpireTime,omitempty"`
	Host               string `json:"Host,omitempty"`
	InstanceId         string `json:"InstanceId,omitempty"`
	InstanceType       string `json:"InstanceType,omitempty"`
	NetworkAccessMode  string `json:"NetworkAccessMode,omitempty"`
	Options            string `json:"Options,omitempty"`
	PassPhrase         string `json:"PassPhrase,omitempty"`
	Password           string `json:"Password,omitempty"`
	Port               int64  `json:"Port,omitempty"`
	Protocol           string `json:"Protocol,omitempty"`
	RegionId           string `json:"RegionId,omitempty"`
	Username           string `json:"Username,omitempty"`
	VpcId              string `json:"VpcId,omitempty"`
}

type LoginInstanceRequestPartnerInfo struct {
	PartnerID   string `json:"PartnerId,omitempty"`
	PartnerName string `json:"PartnerName,omitempty"`
}

type LoginInstanceRequestUserAccountOptions struct {
	LoginLimit string `json:"LoginLimit,omitempty"`
}

type LoginInstanceRequestUserAccount struct {
	AccountID        string `json:"AccountId,omitempty"`
	AccountPlatForm  string `json:"AccountPlatForm,omitempty"`
	AccountStructure string `json:"AccountStructure,omitempty"`
	DurationSeconds  string `json:"DurationSeconds,omitempty"`
	EmpID            string `json:"EmpId,omitempty"`
	ExpireTime       string `json:"ExpireTime,omitempty"`
	LoginName        string `json:"LoginName,omitempty"`
	Options          string `json:"Options,omitempty"`
	ParentID         string `json:"ParentId,omitempty"`
}
