package authenticate

import (
	"github.com/Appkube-awsx/awsx-common/authenticate/aws"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/spf13/cobra"
)

func AuthenticateCommand(cmd *cobra.Command) (bool, *model.Auth, error) {
	cloudElementId, _ := cmd.PersistentFlags().GetString("elementId")
	cloudElementApiUrl, _ := cmd.PersistentFlags().GetString("cmdbApiUrl")
	vaultUrl, _ := cmd.PersistentFlags().GetString("vaultUrl")
	vaultToken, _ := cmd.PersistentFlags().GetString("vaultToken")
	vaultKey, _ := cmd.PersistentFlags().GetString("vaultKey")
	region, _ := cmd.PersistentFlags().GetString("zone")
	acKey, _ := cmd.PersistentFlags().GetString("accessKey")
	secKey, _ := cmd.PersistentFlags().GetString("secretKey")
	crossAccountRoleArn, _ := cmd.PersistentFlags().GetString("crossAccountRoleArn")
	externalId, _ := cmd.PersistentFlags().GetString("externalId")
	commandParam := model.CommandParam{
		CloudElementId:      cloudElementId,
		CloudElementApiUrl:  cloudElementApiUrl,
		VaultUrl:            vaultUrl,
		VaultToken:          vaultToken,
		VaultKey:            vaultKey,
		Region:              region,
		AccessKey:           acKey,
		SecretKey:           secKey,
		CrossAccountRoleArn: crossAccountRoleArn,
		ExternalId:          externalId,
	}
	return DoAuthenticate(commandParam)
}

func AuthenticateSubCommand(cmd *cobra.Command) (bool, *model.Auth, error) {
	cloudElementId, _ := cmd.Parent().PersistentFlags().GetString("elementId")
	cloudElementApiUrl, _ := cmd.Parent().PersistentFlags().GetString("cmdbApiUrl")
	vaultUrl, _ := cmd.Parent().PersistentFlags().GetString("vaultUrl")
	vaultToken, _ := cmd.Parent().PersistentFlags().GetString("vaultToken")
	vaultKey, _ := cmd.Parent().PersistentFlags().GetString("vaultKey")
	region, _ := cmd.Parent().PersistentFlags().GetString("zone")
	acKey, _ := cmd.Parent().PersistentFlags().GetString("accessKey")
	secKey, _ := cmd.Parent().PersistentFlags().GetString("secretKey")
	crossAccountRoleArn, _ := cmd.Parent().PersistentFlags().GetString("crossAccountRoleArn")
	externalId, _ := cmd.Parent().PersistentFlags().GetString("externalId")
	commandParam := model.CommandParam{
		CloudElementId:      cloudElementId,
		CloudElementApiUrl:  cloudElementApiUrl,
		VaultUrl:            vaultUrl,
		VaultToken:          vaultToken,
		VaultKey:            vaultKey,
		Region:              region,
		AccessKey:           acKey,
		SecretKey:           secKey,
		CrossAccountRoleArn: crossAccountRoleArn,
		ExternalId:          externalId,
	}
	return DoAuthenticate(commandParam)
}

func DoAuthenticate(commandParam model.CommandParam) (bool, *model.Auth, error) {
	return aws.GetAwsCreds(commandParam)
}
