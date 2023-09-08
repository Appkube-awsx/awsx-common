package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// Auth struct to store authentication data
type Auth struct {
	VaultUrl            string `json:"vaultUrl,omitempty"`
	VaultToken          string `json:"vaultToken,omitempty"`
	VaultKey            string `json:"vaultKey,omitempty"`
	Region              string `json:"region,omitempty"`
	CrossAccountRoleArn string `json:"crossAccountRoleArn,omitempty"`
	AccessKey           string `json:"accessKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	ExternalId          string `json:"externalId,omitempty"`
}

const (
	LAMBDA_CLIENT         = "lambda"
	COST_EXPLORER         = "costExplorer"
	CLOUDWATCH_LOG        = "cloudWatchLog"
	KMS_CLIENT            = "kms"
	ELBV2_CLIENT          = "elbv2"
	CONFIG_SERVICE_CLIENT = "configservice"
	EKS_CLIENT            = "eks"
)

var clients = map[string]func(*session.Session) interface{}{
	LAMBDA_CLIENT:         func(session *session.Session) interface{} { return lambda.New(session) },
	COST_EXPLORER:         func(session *session.Session) interface{} { return costexplorer.New(session) },
	CLOUDWATCH_LOG:        func(session *session.Session) interface{} { return cloudwatchlogs.New(session) },
	KMS_CLIENT:            func(session *session.Session) interface{} { return kms.New(session) },
	ELBV2_CLIENT:          func(session *session.Session) interface{} { return elbv2.New(session) },
	CONFIG_SERVICE_CLIENT: func(session *session.Session) interface{} { return configservice.New(session) },
	EKS_CLIENT:            func(session *session.Session) interface{} { return eks.New(session) },
}

// GetClient is returns aws clients
func GetClient(auth Auth, clientType string) interface{} {

	// Get session
	awsSession := GetSessionWithAssumeRole(auth)
	return clients[clientType](awsSession)
}
