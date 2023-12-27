package cmdb

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/config"
	"github.com/Appkube-awsx/awsx-common/httpclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"net/http"
)

func GetCloudElement(commandParam model.CommandParam) (*model.CloudElement, error) {
	if commandParam.CloudElementApiUrl != "" {
		return GetCloudElementData(commandParam.CloudElementApiUrl, commandParam.CloudElementId)
	}
	return GetCloudElementData(config.CmdbUrl, commandParam.CloudElementId)
}

func GetCloudElementData(url string, id string) (*model.CloudElement, error) {
	cmdbResp, _, err := httpclient.ExecuteApi(http.MethodGet, url+"/cloud-element/search?id="+id, "", nil)
	if err != nil {
		return nil, fmt.Errorf("cmdb api failed to get cloud element details", err)
	}
	cmdbRespString := string(cmdbResp)
	var out []*model.CloudElement
	err = json.Unmarshal([]byte(cmdbRespString), &out)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error to unmarshal cmdb cloud element response", err)
	}
	if len(out) > 0 {
		return out[0], err
	}
	return nil, fmt.Errorf("no response received from cmdb cloud element api")
}
