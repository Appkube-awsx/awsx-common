package aws

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/cmdb"
	"github.com/Appkube-awsx/awsx-common/config"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/Appkube-awsx/awsx-common/vault"
	"log"
	"os"
)

func GetClientAwsCreds(commandParam model.CommandParam) (*model.ClientAwsSecrets, error) {
	resp, err := getClientAwsCredsFromVault(commandParam)
	if err == nil {
		return resp, err
	}
	resp, err = getClientAwsCredsFromEnvironmentVariable()
	if err == nil {
		return resp, err
	}
	resp, err = getClientAwsSecretsFromCmdb(commandParam)
	if err == nil {
		return resp, err
	}
	resp, err = getClientAwsCredsFromCommandLine(commandParam)
	if err == nil {
		return resp, err
	}
	return nil, err
}

func getClientAwsCredsFromVault(commandParam model.CommandParam) (*model.ClientAwsSecrets, error) {
	log.Println("getting client aws credentials from vault")
	url := commandParam.VaultUrl
	tokan := commandParam.VaultToken
	if commandParam.VaultUrl == "" {
		url = config.VaultUrl
		tokan = config.VaultToken
	}
	if commandParam.VaultKey == "" {
		msg := "vault key not provided to get client aws credentials"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}
	vaultResp, err := vault.GetAccountDetails(url, tokan, commandParam.VaultKey)
	if err != nil {
		log.Println("vault api to get client aws credentials failed: ", err)
		return nil, err
	}
	if vaultResp.Data.ExternalId == "" || vaultResp.Data.CrossAccountRoleArn == "" {
		msg := "client aws credentials not found in vault"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}

	clientAwsSecrets := model.ClientAwsSecrets{
		ExternalId:          vaultResp.Data.ExternalId,
		CrossAccountRoleArn: vaultResp.Data.CrossAccountRoleArn,
	}
	return &clientAwsSecrets, nil
}

func getClientAwsCredsFromEnvironmentVariable() (*model.ClientAwsSecrets, error) {
	log.Println("getting client aws credentials from environment variable")

	clientAwsExternalId := os.Getenv("CLIENT_AWS_EXTERNAL_ID")
	clientAwsCrollAccountRoleArn := os.Getenv("CLIENT_AWS_CROSS_ACCOUNT_ROLE_ARN")

	if clientAwsExternalId == "" || clientAwsCrollAccountRoleArn == "" {
		msg := "client aws credentials not found in environment variable"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}

	clientAwsSecrets := model.ClientAwsSecrets{
		ExternalId:          clientAwsExternalId,
		CrossAccountRoleArn: clientAwsCrollAccountRoleArn,
	}
	return &clientAwsSecrets, nil
}

func getClientAwsSecretsFromCmdb(commandParam model.CommandParam) (*model.ClientAwsSecrets, error) {
	log.Println("getting client aws credentials from cmdb")
	cloudElementResp, err := cmdb.GetCloudElement(commandParam)
	if err != nil {
		return nil, err
	}
	landingZoneResp, err := cmdb.GetLandingZone(commandParam, int(cloudElementResp.LandingzoneId))
	if err != nil {
		return nil, fmt.Errorf("cmdb api failed to get landing-zone response", err)
	}
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error to unmarshal cmdb landing-zone response", err)
	}

	clientAwsSecrets := model.ClientAwsSecrets{
		ExternalId:          landingZoneResp.ExternalId,
		CrossAccountRoleArn: landingZoneResp.RoleArn,
	}
	return &clientAwsSecrets, nil
}

func getClientAwsCredsFromCommandLine(commandParam model.CommandParam) (*model.ClientAwsSecrets, error) {
	log.Println("getting client aws credentials supplied from command-line")
	if commandParam.ExternalId == "" {
		msg := "client aws external id not provided from command-line"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}
	if commandParam.CrossAccountRoleArn == "" {
		msg := "client aws cross account role arn not provided from command-line"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}

	clientAwsSecrets := model.ClientAwsSecrets{
		ExternalId:          commandParam.ExternalId,
		CrossAccountRoleArn: commandParam.CrossAccountRoleArn,
	}
	return &clientAwsSecrets, nil

}
