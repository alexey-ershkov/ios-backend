package v1

import (
	"encoding/json"
	"ios-backend/src/CoinBaseApiRequests/v1/models"
	"ios-backend/src/configs"
	"net/http"
	"os"
	"strings"
)

const (
	mapQuery  = "/v1/cryptocurrency/map?sort=cmc_rank&limit=199"
	protocol  = "https://"
	apiEnvURL = "COINMARKET_URL"
	apiEnvKey = "COINMARKET_API_KEY"
)

type CurrencyApi interface {
	GetMetadata() ([]models.CurrencyMeta, error)
}

type CmcApi struct {
	BaseUrl string
	ApiKey  string
}

func NewCurrencyApi() (CurrencyApi, error) {
	key, exists := os.LookupEnv(apiEnvKey)
	if !exists {
		return nil, configs.NoEnvVarError
	}

	baseUrl, exists := os.LookupEnv(apiEnvURL)
	if !exists {
		return nil, configs.NoEnvVarError
	}
	return CmcApi{
		BaseUrl: baseUrl,
		ApiKey:  key,
	}, nil
}

func (cmc CmcApi) GetMetadata() ([]models.CurrencyMeta, error) {
	rawMapReq, err := cmc.doRequest("/v1/cryptocurrency/map?sort=cmc_rank&limit=200")
	if err != nil {
		return nil, err
	}

	metaData := make([]models.CurrencyMeta, 0)
	symbols := ""

	for _, element := range rawMapReq.([]interface{}) {
		meta := models.CurrencyMeta{}
		jsonElement, err := json.Marshal(element)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonElement, &meta)
		if err != nil {
			return nil, err
		}
		metaData = append(metaData, meta)
		symbols += meta.Symbol + ","
	}

	symbols = strings.TrimSuffix(symbols, ",")

	rawInfoReq, err := cmc.doRequest("/v1/cryptocurrency/info?symbol="+symbols)
	if err != nil {
		return nil, err
	}

	for idx, meta := range metaData {
		currInfo := rawInfoReq.(map[string]interface{})[meta.Symbol]
		jsonCurrInfo, err := json.Marshal(currInfo)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(jsonCurrInfo, &metaData[idx])
		if err != nil {
			return nil, err
		}
	}

	return metaData, nil
}

func (cmc CmcApi) doRequest(query string) (interface{}, error) {
	client := http.Client{}
	mapReq, err := http.NewRequest("GET", cmc.BaseUrl+query, nil)
	if err != nil {
		return nil, err
	}

	mapReq.Header.Add("X-CMC_PRO_API_KEY", cmc.ApiKey)

	resp, err := client.Do(mapReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	data = data.(map[string]interface{})["data"]
	return data, nil
}
