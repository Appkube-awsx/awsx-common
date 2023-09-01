package vault

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AwsCredential struct {
	Region              string `json:"region,omitempty"`
	AccessKey           string `json:"accessKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	CrossAccountRoleArn string `json:"crossAccountRoleArn,omitempty"`
	ExternalId          string `json:"externalId,omitempty"`
}

type VaultResponse struct {
	RequestId     string        `json:"request_id,omitempty"`
	LeaseId       string        `json:"lease_id,omitempty"`
	Renewable     string        `json:"renewable,omitempty"`
	LeaseDuration int64         `json:"lease_duration,omitempty"`
	Data          AwsCredential `json:"data,omitempty"`
}

func GetAccountDetails(vaultUrl string, vaultToken string, accountNo string) (*VaultResponse, error) {
	log.Println("Calling account details API")
	client := &http.Client{}
	req, err := http.NewRequest("GET", vaultUrl+"/"+accountNo, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Vault-Token", vaultToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var responseObject VaultResponse
	json.Unmarshal(bodyBytes, &responseObject)
	return &responseObject, nil

}
