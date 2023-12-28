package aws

import (
	"github.com/Appkube-awsx/awsx-common/model"
)

func GetAwsCreds(commandParam model.CommandParam) (bool, *model.Auth, error) {
	globalAwsCreds, err := GetGlobalAwsCreds(commandParam)
	if err != nil {
		return false, nil, err
	}
	clientAwsCreds, err := GetClientAwsCreds(commandParam)
	if err != nil {
		return false, nil, err
	}

	clientAuth := model.Auth{
		AccessKey:           globalAwsCreds.AccessKey,
		SecretKey:           globalAwsCreds.SecretKey,
		Region:              globalAwsCreds.Region,
		ExternalId:          clientAwsCreds.ExternalId,
		CrossAccountRoleArn: clientAwsCreds.CrossAccountRoleArn,
	}
	return true, &clientAuth, nil
}
