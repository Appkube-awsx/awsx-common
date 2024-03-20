package model

type GlobalAwsSecrets struct {
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
	Region    string `json:"region,omitempty"`
}
type ClientAwsSecrets struct {
	ExternalId          string `json:"externalId,omitempty"`
	CrossAccountRoleArn string `json:"crossAccountRoleArn,omitempty"`
}

type CloudElement struct {
	Id                       int64                  `json:"id"`
	ElementType              string                 `json:"elementType,omitempty"`
	HostedServices           map[string]interface{} `json:"hostedServices,omitempty"`
	Arn                      string                 `json:"arn,omitempty"`
	InstanceId               string                 `json:"instanceId,omitempty"`
	InstanceName             string                 `json:"instanceName,omitempty"`
	Category                 string                 `json:"category,omitempty"`
	SlaJson                  map[string]interface{} `json:"slaJson,omitempty"`
	CostJson                 map[string]interface{} `json:"costJson,omitempty"`
	ViewJson                 map[string]interface{} `json:"viewJson,omitempty"`
	ConfigJson               map[string]interface{} `json:"configJson,omitempty"`
	ComplianceJson           map[string]interface{} `json:"complianceJson,omitempty"`
	Status                   string                 `json:"status,omitempty"`
	CreatedBy                string                 `json:"createdBy,omitempty"`
	UpdatedBy                string                 `json:"updatedBy,omitempty"`
	CreatedOn                string                 `json:"createdOn,omitempty"`
	UpdatedOn                string                 `json:"updatedOn,omitempty"`
	LogGroupName             string                 `json:"logGroupName,omitempty"`
	LandingzoneId            int64                  `json:"landingzoneId"`
	LandingZone              string                 `json:"landingZone,omitempty"`
	DbCategoryId             int64                  `json:"dbCategoryId"`
	DbCategoryName           string                 `json:"dbCategoryName,omitempty"`
	ProductEnclaveId         int64                  `json:"productEnclaveId"`
	ProductEnclaveInstanceId string                 `json:"productEnclaveInstanceId,omitempty"`
}

type Landingzone struct {
	Id               int64  `json:"id"`
	Description      string `json:"description,omitempty"`
	LandingZone      string `json:"landingZone,omitempty"`
	Cloud            string `json:"cloud,omitempty"`
	DisplayName      string `json:"displayName,omitempty"`
	RoleArn          string `json:"roleArn,omitempty"`
	ExternalId       string `json:"externalId,omitempty"`
	Status           string `json:"status,omitempty"`
	CreatedBy        string `json:"createdBy,omitempty"`
	UpdatedBy        string `json:"updatedBy,omitempty"`
	CreatedOn        string `json:"createdOn,omitempty"`
	UpdatedOn        string `json:"updatedOn,omitempty"`
	DepartmentId     int64  `json:"departmentId"`
	DepartmentName   string `json:"departmentName,omitempty"`
	OrganizationId   int64  `json:"organizationId"`
	OrganizationName string `json:"organizationName,omitempty"`
}

type ApiResponse struct {
	Status     string        `json:"status,omitempty"`
	Message    string        `json:"message,omitempty"`
	StatusCode int64         `json:"statusCode,omitempty"`
	Data       AwsCredential `json:"data,omitempty"`
}

type AwsCredential struct {
	Region              string `json:"region,omitempty"`
	AccessKey           string `json:"accessKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	CrossAccountRoleArn string `json:"crossAccountRoleArn,omitempty"`
	ExternalId          string `json:"externalId,omitempty"`
}

type VaultResponse struct {
	RequestId     string        `json:"request_id,omitempty"`
	LeaseId       string        `json:"lease_id,omitempty"`
	Renewable     string        `json:"renewable,omitempty"`
	LeaseDuration int64         `json:"lease_duration,omitempty"`
	Data          AwsCredential `json:"data,omitempty"`
}

type CommandParam struct {
	LandingZoneId       string `json:"landingZoneId,omitempty"`
	CloudElementId      string `json:"cloudElementId,omitempty"`
	CloudElementApiUrl  string `json:"cloudElementApiUrl,omitempty"`
	VaultUrl            string `json:"vaultUrl,omitempty"`
	VaultToken          string `json:"vaultToken,omitempty"`
	VaultKey            string `json:"vaultKey,omitempty"`
	Region              string `json:"region,omitempty"`
	AccessKey           string `json:"accessKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	CrossAccountRoleArn string `json:"crossAccountRoleArn,omitempty"`
	ExternalId          string `json:"externalId,omitempty"`
}

// Auth struct to store authentication data
type Auth struct {
	LandingZoneId       string `json:"landingZoneId,omitempty"`
	VaultUrl            string `json:"vaultUrl,omitempty"`
	VaultToken          string `json:"vaultToken,omitempty"`
	VaultKey            string `json:"vaultKey,omitempty"`
	Region              string `json:"region,omitempty"`
	CrossAccountRoleArn string `json:"crossAccountRoleArn,omitempty"`
	AccessKey           string `json:"accessKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	ExternalId          string `json:"externalId,omitempty"`
}
