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
		"us-east-1",
		os.Getenv("AWS_CROSS_ARN"),
		os.Getenv("AWS_ACCKEY"),
		os.Getenv("AWS_SECKEY"),
		os.Getenv("AWS_EXTERNALID"),
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
