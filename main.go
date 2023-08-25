package main

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/lambda"
	"os"
)

// It is just for testing
func main() {

	auth := client.Auth{
		Region:              "us-east-1",
		CrossAccountRoleArn: os.Getenv("AWS_CROSS_ARN"),
		AccessKey:           os.Getenv("AWS_ACCKEY"),
		SecretKey:           os.Getenv("AWS_SECKEY"),
		ExternalId:          os.Getenv("AWS_EXTERNALID"),
	}

	data := client.GetClient(auth, client.LAMBDA_CLIENT)

	a := data.(*lambda.Lambda)

	input := lambda.ListFunctionsInput{}

	result, err := a.ListFunctions(&input)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("this is returned client", result)

}
