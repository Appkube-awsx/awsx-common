package client

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/awssession"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
)

func GetSessionWithAssumeRole(roleArn string, sessionName string, externalId string, accesskey string, secretKey string, region string) *session.Session {
	sess, err := awssession.GetSessionByCreds(region, accesskey, secretKey, "")

	if err != nil {
		fmt.Printf("failed to create aws session, %v\n", err)
		log.Fatal(err)
	}

	svc := sts.New(sess)

	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
	}

	if externalId != "nil" {
		fmt.Println("Trying to fetch external id to assume new role")
		assumeRoleInput.ExternalId = aws.String(externalId)
	}

	result, err := svc.AssumeRole(&assumeRoleInput)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}

	awsSession, err := awssession.GetSessionByCreds(region, *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}

	return awsSession
}
