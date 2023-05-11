package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// Auth struct to store authentication data
type Auth struct {
	Region              string
	CrossAccountRoleArn string
	AccessKey           string
	SecretKey           string
	ExternalId          string
}

const (
	LAMBDA_CLIENT  = "lambda"
	COST_EXPLORER  = "costExplorer"
	CLOUDWATCH_LOG = "cloudWatchLog"
)

var clients = map[string]func(*session.Session) interface{}{
	LAMBDA_CLIENT:  func(session *session.Session) interface{} { return lambda.New(session) },
	COST_EXPLORER:  func(session *session.Session) interface{} { return costexplorer.New(session) },
	CLOUDWATCH_LOG: func(session *session.Session) interface{} { return cloudwatch.New(session) },
}

// GetClient is returns aws clients
func GetClient(auth Auth, clientType string) interface{} {

	// Get session
	awsSession := GetSessionWithAssumeRole(auth)
	return clients[clientType](awsSession)
}
