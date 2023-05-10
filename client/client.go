package client

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/costexplorer"
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

// GetClient is returns aws clients
func GetClient(auth Auth, clientType string) interface{} {

	// Get session
	awsSession := GetSessionWithAssumeRole(auth.CrossAccountRoleArn, sessionName, auth.ExternalId, auth.AccessKey, auth.SecretKey, auth.Region)

	switch clientType {

	case "lambda":
		return lambda.New(awsSession)
	case "costExplorer":
		return costexplorer.New(awsSession)
	case "cloudWatchLog":
		return cloudwatchlogs.New(awsSession)

	}

	return nil
}
