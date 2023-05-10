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

// GetLambdaClient is returns aws lambda client
func GetLambdaClient(auth Auth) *lambda.Lambda {

	// Get session
	awsSession := GetSessionWithAssumeRole(auth.CrossAccountRoleArn, sessionName, auth.ExternalId, auth.AccessKey, auth.SecretKey, auth.Region)

	// Get client
	lambdaClient := lambda.New(awsSession)

	// return client
	return lambdaClient
}

// GetCostClient is returns aws costexplorer client
func GetCostClient(auth Auth) *costexplorer.CostExplorer {

	// Get session
	awsSession := GetSessionWithAssumeRole(auth.CrossAccountRoleArn, sessionName, auth.ExternalId, auth.AccessKey, auth.SecretKey, auth.Region)

	// Get client
	costexplorer := costexplorer.New(awsSession)

	// return client
	return costexplorer
}

// GetCloudWatchLogClient is returns aws cloudwatchlogs client
func GetCloudWatchLogClient(auth Auth) *cloudwatchlogs.CloudWatchLogs {

	// Get session
	awsSession := GetSessionWithAssumeRole(auth.CrossAccountRoleArn, sessionName, auth.ExternalId, auth.AccessKey, auth.SecretKey, auth.Region)

	// Get client
	cloudwatchlogs := cloudwatchlogs.New(awsSession)

	// return client
	return cloudwatchlogs
}
