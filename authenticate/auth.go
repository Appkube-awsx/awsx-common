package authenticate

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-common/vault"
	"github.com/spf13/cobra"
	"log"
)

func AuthenticateData(vaultUrl string, vaultToken string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) (bool, *client.Auth, error) {
	if vaultUrl != "" {
		log.Println("vault url provided. getting user credentials from vault")
		if vaultToken == "" {
			log.Println("vault token missing")
			return false, nil, fmt.Errorf("vault token missing")
		}
		if accountNo == "" {
			log.Println("account no missing")
			return false, nil, fmt.Errorf("account no missing")
		}
		log.Println("Getting account details from vault")
		vaultResp, err := vault.GetAccountDetails(vaultUrl, vaultToken, accountNo)
		if err != nil {
			log.Println("Error in calling vault api to get account details. \n", err)
			return false, nil, err
		}
		if vaultResp.Data.AccessKey == "" || vaultResp.Data.SecretKey == "" || vaultResp.Data.CrossAccountRoleArn == "" || vaultResp.Data.ExternalId == "" {
			log.Println("account details not found in vault")
			return false, nil, fmt.Errorf("account details not found in vault")
		}

		clientAuth := client.Auth{
			CrossAccountRoleArn: vaultResp.Data.CrossAccountRoleArn,
			AccessKey:           vaultResp.Data.AccessKey,
			SecretKey:           vaultResp.Data.SecretKey,
			ExternalId:          vaultResp.Data.ExternalId,
		}
		if region != "" {
			clientAuth.Region = region
		} else {
			log.Println("region not provided. default region will be used")
			clientAuth.Region = vaultResp.Data.Region
		}
		return true, &clientAuth, nil
	}
	log.Println("vault url not provided. validating provided user credentials")
	if region == "" {
		log.Println("region missing")
		return false, nil, fmt.Errorf("region missing")
	}
	if acKey == "" {
		log.Println("access key missing")
		return false, nil, fmt.Errorf("access key missing")
	}
	if secKey == "" {
		log.Println("secret key missing")
		return false, nil, fmt.Errorf("secret key missing")
	}
	if crossAccountRoleArn == "" {
		log.Println("cross account role arn missing")
		return false, nil, fmt.Errorf("cross account role arn missing")
	}
	if externalId == "" {
		log.Println("external id missing")
		return false, nil, fmt.Errorf("external id missing")
	}
	clientAuth := client.Auth{
		Region:              region,
		CrossAccountRoleArn: crossAccountRoleArn,
		AccessKey:           acKey,
		SecretKey:           secKey,
		ExternalId:          externalId,
	}
	return true, &clientAuth, nil
}

func CommandAuth(cmd *cobra.Command) (bool, *client.Auth, error) {
	vaultUrl, _ := cmd.PersistentFlags().GetString("vaultUrl")
	vaultToken, _ := cmd.PersistentFlags().GetString("vaultToken")
	accountNo, _ := cmd.PersistentFlags().GetString("accountId")
	region, _ := cmd.PersistentFlags().GetString("zone")
	acKey, _ := cmd.PersistentFlags().GetString("accessKey")
	secKey, _ := cmd.PersistentFlags().GetString("secretKey")
	crossAccountRoleArn, _ := cmd.Flags().GetString("crossAccountRoleArn")
	externalId, _ := cmd.PersistentFlags().GetString("externalId")
	authFlag, clientAuth, err := AuthenticateData(vaultUrl, vaultToken, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
	return authFlag, clientAuth, err
}
