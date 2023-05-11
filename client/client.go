package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var sessionName = "assume_role_session_name"

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
	COST_EXPLORER:  func(session *session.Session) interface{} { return lambda.New(session) },
	CLOUDWATCH_LOG: func(session *session.Session) interface{} { return lambda.New(session) },
}

// GetClient is returns aws clients
func GetClient(auth Auth, clientType string) interface{} {

	// Get session
	awsSession := GetSessionWithAssumeRole(auth.CrossAccountRoleArn, sessionName, auth.ExternalId, auth.AccessKey, auth.SecretKey, auth.Region)
	return clients[clientType](awsSession)
}
