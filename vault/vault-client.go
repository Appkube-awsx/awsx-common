package vault

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/httpclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"log"
	"net/http"
)

func GetAccountDetails(vaultUrl string, vaultToken string, vaultKey string) (*model.VaultResponse, error) {
	log.Println("getting aws credentials from vault")
	headers := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"X-Vault-Token": vaultToken,
	}
	vaultResp, _, err := httpclient.ExecuteApi(http.MethodGet, vaultUrl+"/"+vaultKey, "", headers)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	responseObject := model.VaultResponse{}
	err = json.Unmarshal(vaultResp, &responseObject)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error to unmarshal vault response", err)
	}
	return &responseObject, nil

}
