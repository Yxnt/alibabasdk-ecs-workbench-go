package models

type LoginInstanceRequest struct {
}

type LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo struct {
}

type LoginInstanceRequestInstanceLoginInfoOptions struct {
}

type LoginInstanceRequestInstanceLoginInfo struct {
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
