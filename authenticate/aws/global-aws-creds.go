package aws

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/config"
	"github.com/Appkube-awsx/awsx-common/httpclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/Appkube-awsx/awsx-common/vault"
	"log"
	"net/http"
	"os"
)

func GetGlobalAwsCreds(commandParam model.CommandParam) (*model.GlobalAwsSecrets, error) {
	resp, err := getGlobalAwsCredsFromCommandLine(commandParam)
	if err == nil {
		return resp, err
	}
	if err == nil {
		return resp, err
	}
	resp, err = getGlobalAwsSecretsFromCmdb(commandParam)
	if err == nil {
		return resp, err
	}
	resp, err = getGlobalAwsCredsFromVault(commandParam)
	if err == nil {
		return resp, err
	}
	resp, err = getGlobalAwsCredsFromEnvironmentVariable(commandParam)
	return nil, err
}

func getGlobalAwsCredsFromVault(commandParam model.CommandParam) (*model.GlobalAwsSecrets, error) {
	log.Println("getting global aws credentials from vault")
	url := commandParam.VaultUrl
	tokan := commandParam.VaultToken
	if commandParam.VaultUrl == "" {
		url = config.VaultUrl
		tokan = config.VaultToken
	}
	vaultResp, err := vault.GetAccountDetails(url, tokan, config.GlobalAwsSecrets)

	if err != nil {
		log.Println("vault api to get global aws credentials failed: ", err)
		return nil, err
	}
	if vaultResp.Data.AccessKey == "" || vaultResp.Data.SecretKey == "" {
		msg := "global aws credentials not found in vault"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}

	region := ""
	if commandParam.Region != "" {
		log.Println("region supplied from command line")
		region = commandParam.Region
	} else if vaultResp.Data.Region != "" {
		log.Println("region supplied from vault")
		region = vaultResp.Data.Region
	} else {
		log.Println("default region us-east-1 being used. it is neither supplied from command line nor found in vault")
		region = "us-east-1"
	}

	globalAwsSecrets := model.GlobalAwsSecrets{
		AccessKey: vaultResp.Data.AccessKey,
		SecretKey: vaultResp.Data.SecretKey,
		Region:    region,
	}
	return &globalAwsSecrets, nil
}

func getGlobalAwsCredsFromEnvironmentVariable(commandParam model.CommandParam) (*model.GlobalAwsSecrets, error) {
	log.Println("getting global aws credentials from environment variable")

	globalAwsAccessKey := os.Getenv("GLOBAL_AWS_ACCESS_KEY")
	globalAwsSecretKey := os.Getenv("GLOBAL_AWS_SECRET_KEY")
	globalAwsRegion := os.Getenv("GLOBAL_AWS_REGION")

	if globalAwsAccessKey == "" || globalAwsSecretKey == "" {
		msg := "global aws credentials not found in environment variable"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}

	region := ""
	if commandParam.Region != "" {
		log.Println("region supplied from command line")
		region = commandParam.Region
	} else if globalAwsRegion != "" {
		log.Println("region supplied from environment variable")
		region = globalAwsRegion
	} else {
		log.Println("default region us-east-1 being used. it is neither supplied from command line nor found in environment variable")
		region = "us-east-1"
	}

	globalAwsSecrets := model.GlobalAwsSecrets{
		AccessKey: globalAwsAccessKey,
		SecretKey: globalAwsSecretKey,
		Region:    region,
	}
	return &globalAwsSecrets, nil
}

func getGlobalAwsSecretsFromCmdb(commandParam model.CommandParam) (*model.GlobalAwsSecrets, error) {
	log.Println("getting global aws credentials from cmdb")
	type cmdbResp struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	}
	acKeyResp, acKeyRespStatusCode, err := httpclient.ExecuteApi(http.MethodGet, config.CmdbUrl+"/config/decrypt/get-by-key/GLOBAL_AWS_ACCESS_KEY", "", nil)
	if err != nil || acKeyRespStatusCode >= 400 {
		return nil, fmt.Errorf("cmdb api failed to get global aws access key", err)
	}

	secKeyResp, secKeyRespStatusCode, err := httpclient.ExecuteApi(http.MethodGet, config.CmdbUrl+"/config/decrypt/get-by-key/GLOBAL_AWS_SECRET_KEY", "", nil)
	if err != nil || secKeyRespStatusCode >= 400 {
		return nil, fmt.Errorf("cmdb api failed to get global aws secret key", err)
	}

	regionResp, regionRespStatusCode, regionErr := httpclient.ExecuteApi(http.MethodGet, config.CmdbUrl+"/config/get-by-key/GLOBAL_AWS_REGION", "", nil)
	regionKey := cmdbResp{}
	isRegionFound := false
	if regionErr == nil && regionRespStatusCode < 400 {
		isRegionFound = true
		//log.Println("global aws region found in cmdb")
		regionRespString := string(regionResp)
		err = json.Unmarshal([]byte(regionRespString), &regionKey)
		if err != nil {
			return nil, fmt.Errorf("json unmarshal error to unmarshal global aws region", err)
		}
	}

	acKeyRespString := string(acKeyResp)
	secKeyRespString := string(secKeyResp)

	acKey := cmdbResp{}
	secKey := cmdbResp{}

	err = json.Unmarshal([]byte(acKeyRespString), &acKey)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error to unmarshal global aws access key", err)
	}
	err = json.Unmarshal([]byte(secKeyRespString), &secKey)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error to unmarshal global aws secret key", err)
	}

	region := ""
	if commandParam.Region != "" {
		log.Println("region supplied from command line")
		region = commandParam.Region
	} else if isRegionFound {
		log.Println("region supplied from cmdb")
		region = regionKey.Value
	} else {
		log.Println("default region us-east-1 being used. it is neither supplied from command line nor from cmdb")
		region = "us-east-1"
	}

	globalAwsSecret := model.GlobalAwsSecrets{
		AccessKey: acKey.Value,
		SecretKey: secKey.Value,
		Region:    region,
	}
	return &globalAwsSecret, nil
}

func getGlobalAwsCredsFromCommandLine(commandParam model.CommandParam) (*model.GlobalAwsSecrets, error) {
	log.Println("getting global aws credentials supplied from command-line")
	if commandParam.AccessKey == "" {
		msg := "global aws access key not provided from command-line"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}
	if commandParam.SecretKey == "" {
		msg := "global aws secret key not provided from command-line"
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}
	region := ""
	if commandParam.Region == "" {
		log.Println("default region us-east-1 being used. it is not supplied from command line")
		region = "us-east-1"
	}

	globalAwsSecrets := model.GlobalAwsSecrets{
		AccessKey: commandParam.AccessKey,
		SecretKey: commandParam.SecretKey,
		Region:    region,
	}
	return &globalAwsSecrets, nil
}
