package main

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate/aws"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/cmdb"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// It is just for testing
func main() {
	Execute()
}

var TestCmd = &cobra.Command{
	Use:   "testCommonFunctions",
	Short: "command to test common functions",
	Long:  `command to test common functions`,
	Run: func(cmd *cobra.Command, args []string) {
		testAwsCreds(cmd)
	},
}

func testAwsCreds(cmd *cobra.Command) {
	commandParam := model.CommandParam{}
	_, resp, err := aws.GetAwsCreds(commandParam)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response : ", resp)
}

func testClientAwsCreds(cmd *cobra.Command) {
	commandParam := model.CommandParam{}
	resp, err := aws.GetClientAwsCreds(commandParam)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response : ", resp)
}

func testGlobalAwsCreds(cmd *cobra.Command) {
	commandParam := model.CommandParam{}
	resp, err := aws.GetGlobalAwsCreds(commandParam)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response : ", resp)
}

func testLandingZone(cmd *cobra.Command) {
	cloudElementApiUrl, _ := cmd.PersistentFlags().GetString("cloudElementApiUrl")
	commandParam := model.CommandParam{
		CloudElementApiUrl: cloudElementApiUrl,
	}
	resp, err := cmdb.GetLandingZone(commandParam, 26)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response : ", resp)
}

func testCloudElement(cmd *cobra.Command) {
	cloudElementId, _ := cmd.PersistentFlags().GetString("cloudElementId")
	//cloudElementApiUrl, _ := cmd.PersistentFlags().GetString("cloudElementApiUrl")
	commandParam := model.CommandParam{
		CloudElementId: cloudElementId,
		//CloudElementApiUrl: cloudElementApiUrl,
	}
	resp, err := cmdb.GetCloudElement(commandParam)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response : ", resp)
}

func Execute() {
	err := TestCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	TestCmd.PersistentFlags().String("cloudElementId", "", "cmdb cloud element id")
	TestCmd.PersistentFlags().String("cloudElementApiUrl", "", "cmdb cloud element api url")
	TestCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	TestCmd.PersistentFlags().String("vaultToken", "", "vault token")
	TestCmd.PersistentFlags().String("accountId", "", "aws account number")
	TestCmd.PersistentFlags().String("zone", "", "aws region")
	TestCmd.PersistentFlags().String("accessKey", "", "aws access key")
	TestCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	TestCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	TestCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}

func testAwsLambda() {
	auth := model.Auth{
		Region:              "us-east-1",
		CrossAccountRoleArn: os.Getenv("AWS_CROSS_ARN"),
		AccessKey:           os.Getenv("AWS_ACCKEY"),
		SecretKey:           os.Getenv("AWS_SECKEY"),
		ExternalId:          os.Getenv("AWS_EXTERNALID"),
	}
	data := awsclient.GetClient(auth, awsclient.LAMBDA_CLIENT)
	a := data.(*lambda.Lambda)
	input := lambda.ListFunctionsInput{}
	result, err := a.ListFunctions(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("this is returned awsclient", result)
}
