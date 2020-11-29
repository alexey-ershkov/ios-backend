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
	apiEnvURL = "COINMARKET_URL"
	apiEnvKey = "COINMARKET_API_KEY"
	fiats     = "FIATS"
)

type CurrencyApi interface {
	GetMetadata() ([]models.CurrencyMeta, error)
	GetFiatMetadata() ([]models.FiatMeta, error)
	GetCurrCurrencyInfo() ([]models.CurrCryptoInfo, error)
}

type CmcApi struct {
	BaseUrl string
	ApiKey  string
	Fiats   []string
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

	fiats, exists := os.LookupEnv(fiats)
	if !exists {
		return nil, configs.NoEnvVarError
	}
	fiats = "USD," + fiats

	return CmcApi{
		BaseUrl: baseUrl,
		ApiKey:  key,
		Fiats:   strings.Split(fiats, ","),
	}, nil
}

func (cmc CmcApi) GetMetadata() ([]models.CurrencyMeta, error) {
	rawMapReq, err := cmc.doRequest("/v1/cryptocurrency/map?sort=cmc_rank&limit=100")
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

	rawInfoReq, err := cmc.doRequest("/v1/cryptocurrency/info?symbol=" + symbols)
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

func (cmc CmcApi) GetFiatMetadata() ([]models.FiatMeta, error) {
	rawFiatMapReq, err := cmc.doRequest("/v1/fiat/map")
	if err != nil {
		return nil, err
	}

	fiatModels := make([]models.FiatMeta, 0)

	for _, fiat := range cmc.Fiats {
		fiatModel := cmc.findFiat(rawFiatMapReq.([]interface{}), fiat)
		if fiatModel == nil {
			return nil, configs.NoSuchFiat
		}
		fiatModels = append(fiatModels, *fiatModel)
	}

	return fiatModels, nil
}

func (cmc CmcApi) GetCurrCurrencyInfo() ([]models.CurrCryptoInfo, error) {
	currCryptoModels := make([]models.CurrCryptoInfo, 0)

	for fiatIdx, fiat := range cmc.Fiats {
		rawCurrCryptoReq, err := cmc.doRequest("/v1/cryptocurrency/listings/" +
			"latest?limit=100&convert=" + fiat)
		if err != nil {
			return nil, err
		}

		for idx, element := range rawCurrCryptoReq.([]interface{}) {
			if fiatIdx == 0 {
				parsedElem := cmc.parseCurrCrypto(element)
				if parsedElem == nil {
					return nil, configs.CurrUpdParseError
				}
				currCryptoModels = append(currCryptoModels, *parsedElem)
			} else {
				err := cmc.upendFiat(&currCryptoModels[idx], element.(map[string]interface{}), fiat)
				if err != nil {
					return nil, err
				}
			}
		}

	}

	return currCryptoModels, nil
}

func (cmc CmcApi) upendFiat(info *models.CurrCryptoInfo, raw map[string]interface{}, fiat string) error {
	fiatModel := &models.Fiat{}
	quote := raw["quote"].(map[string]interface{})[fiat]

	jsonElement, err := json.Marshal(quote)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonElement, fiatModel)
	if err != nil {
		return err
	}
	fiatModel.Symbol = fiat
	fiatModel.CmcId = info.CmcId
	info.CostInFiats = append(info.CostInFiats, *fiatModel)

	return nil
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

func (cmc CmcApi) findFiat(slice []interface{}, val string) *models.FiatMeta {
	out := &models.FiatMeta{}

	for _, item := range slice {
		jsonElement, err := json.Marshal(item)
		if err != nil {
			return nil
		}
		err = json.Unmarshal(jsonElement, out)
		if err != nil {
			return nil
		}
		if out.Symbol == val {
			return out
		}
	}
	return nil
}

func (cmc CmcApi) parseCurrCrypto(raw interface{}) *models.CurrCryptoInfo {
	out := &models.CurrCryptoInfo{}
	jsonElement, err := json.Marshal(raw)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(jsonElement, out)
	if err != nil {
		return nil
	}

	usdQuote := raw.(map[string]interface{})["quote"].(map[string]interface{})["USD"]

	out.PercentChange1h = usdQuote.(map[string]interface{})["percent_change_1h"].(float64)
	out.PercentChange24h = usdQuote.(map[string]interface{})["percent_change_24h"].(float64)
	out.PercentChange7d = usdQuote.(map[string]interface{})["percent_change_7d"].(float64)

	fiatModel := &models.Fiat{}
	jsonElement, err = json.Marshal(usdQuote)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(jsonElement, fiatModel)
	if err != nil {
		return nil
	}
	fiatModel.Symbol = "USD"
	fiatModel.CmcId = out.CmcId
	out.CostInFiats = append(out.CostInFiats, *fiatModel)

	return out
}
