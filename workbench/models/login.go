package models

type LoginInstanceRequest struct {
	InstanceLoginInfo LoginInstanceRequestInstanceLoginInfo `json:"InstanceLoginInfo"`
	PartnerInfo       LoginInstanceRequestPartnerInfo       `json:"PartnerInfo"`
	UserAccount       LoginInstanceRequestUserAccount       `json:"UserAccount"`
}

type LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo struct {
	ClusterId     string `json:"ClusterId"`
	ClusterName   string `json:"ClusterName"`
	ContainerName string `json:"ContainerName"`
	Deployment    string `json:"Deployment"`
	Endpoint      string `json:"Endpoint"`
	Headers       string `json:"Headers"`
	Namespace     string `json:"Namespace"`
	PodName       string `json:"PodName"`
}

type LoginInstanceRequestInstanceLoginInfoOptions struct {
	AudioMuteSeconds                 int64                                                     `json:"AudioMuteSeconds"`
	ContainerInfo                    LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo `json:"ContainerInfo"`
	NotificationEventTypes           string                                                    `json:"NotificationEventTypes"`
	NotificationRecipientUrl         string                                                    `json:"NotificationRecipientUrl"`
	NotificationRetryIntervalSeconds string                                                    `json:"NotificationRetryIntervalSeconds"`
	NotificationRetryLimit           string                                                    `json:"NotificationRetryLimit"`
	OperationDisableSeconds          int64                                                     `json:"OperationDisableSeconds"`
	SessionControl                   string                                                    `json:"SessionControl"`
	VideoFreezeSeconds               string                                                    `json:"VideoFreezeSeconds"`
}

type LoginInstanceRequestInstanceLoginInfo struct {
	AuthenticationType string `json:"AuthenticationType"`
	Certificate        string `json:"Certificate"`
	DurationSeconds    string `json:"DurationSeconds"`
	ExpireTime         string `json:"ExpireTime"`
	Host               string `json:"Host"`
	InstanceId         string `json:"InstanceId"`
	InstanceType       string `json:"InstanceType"`
	NetworkAccessMode  string `json:"NetworkAccessMode"`
	Options            string `json:"Options"`
	PassPhrase         string `json:"PassPhrase"`
	Password           string `json:"Password"`
	Port               int64  `json:"Port"`
	Protocol           string `json:"Protocol"`
	RegionId           string `json:"RegionId"`
	Username           string `json:"Username"`
	VpcId              string `json:"VpcId"`
}

type LoginInstanceRequestPartnerInfo struct {
	PartnerID   string `json:"PartnerId"`
	PartnerName string `json:"PartnerName"`
}

type LoginInstanceRequestUserAccountOptions struct {
	LoginLimit string `json:"LoginLimit"`
}

type LoginInstanceRequestUserAccount struct {
	AccountID        string `json:"AccountId"`
	AccountPlatForm  string `json:"AccountPlatForm"`
	AccountStructure string `json:"AccountStructure"`
	DurationSeconds  string `json:"DurationSeconds"`
	EmpID            string `json:"EmpId"`
	ExpireTime       string `json:"ExpireTime"`
	LoginName        string `json:"LoginName"`
	Options          string `json:"Options"`
	ParentID         string `json:"ParentId"`
}
