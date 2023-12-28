package cmdb

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/config"
	"github.com/Appkube-awsx/awsx-common/httpclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"net/http"
	"strconv"
)

func GetLandingZone(commandParam model.CommandParam, landingZoneId int) (*model.Landingzone, error) {
	if commandParam.CloudElementApiUrl != "" {
		return GetLandingZoneData(commandParam.CloudElementApiUrl, landingZoneId)
	}
	return GetLandingZoneData(config.CmdbUrl, landingZoneId)
}

func GetLandingZoneData(url string, landingZoneId int) (*model.Landingzone, error) {
	landingZoneResp, _, err := httpclient.ExecuteApi(http.MethodGet, url+"/landingzone/"+strconv.Itoa(landingZoneId), "", nil)
	if err != nil {
		return nil, fmt.Errorf("cmdb api failed to get landing-zone response", err)
	}
	landingZoneRespString := string(landingZoneResp)
	landingZone := model.Landingzone{}
	err = json.Unmarshal([]byte(landingZoneRespString), &landingZone)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error to unmarshal cmdb landing-zone response", err)
	}
	return &landingZone, nil
}
