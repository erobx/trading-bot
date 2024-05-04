package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/erobx/trading-bot/pkg/app/model"
)

var (
	cs2AppId = "730"
)

func getAssetClassInfo(class_count string, classIds []string, key string) {
	req := model.NewRequest().
		WithGenInterface("ISteamEconomy").
		WithMethod("GetAssetClassInfo").
		WithApiKey(key).
		WithAppId(cs2AppId)

	reqUrl := req.CombineUrl() + "class_count=" + class_count + "&"
	for i, id := range classIds {
		reqUrl = reqUrl + "classid" + strconv.Itoa(i) + "=" + id
		if i != len(classIds)-1 {
			reqUrl = reqUrl + "&"
		}
	}

	_, err := ParseAssetClassInfo(reqUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func getAssetPrices(key string) {
	req := model.NewRequest().
		WithGenInterface("ISteamEconomy").
		WithMethod("GetAssetPrices").
		WithApiKey(key).
		WithAppId(cs2AppId)

	reqUrl := req.CombineUrl() + "languange=en&currency=USD"

	ap, err := ParseAssetPrices(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	writeAssetPrices(ap)
}

func writeAssetPrices(ap AssetPrices) {
	jsonData, err := json.MarshalIndent(ap, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("asset_prices.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

type AssetClassInfo struct {
}

func ParseAssetClassInfo(url string) (AssetClassInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	test := string(body)
	fmt.Println(test)

	aci := AssetClassInfo{}
	// if err := json.Unmarshal(body, &aci); err != nil {
	// 	return AssetClassInfo{}, err
	// }

	return aci, nil
}

type AssetPrices struct {
	Result struct {
		Success bool `json:"success"`
		Assets  []struct {
			Prices struct {
				USD int `json:"USD"`
			} `json:"prices"`
			Name  string `json:"name"`
			Date  string `json:"data"`
			Class []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"class"`
			Classid        string `json:"classid"`
			OriginalPrices struct {
				USD int `json:"USD"`
			} `json:"original_prices,omitempty"`
		} `json:"assets"`
	} `json:"result"`
}

func ParseAssetPrices(url string) (AssetPrices, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ap := AssetPrices{}
	if err := json.Unmarshal(body, &ap); err != nil {
		return AssetPrices{}, err
	}

	return ap, nil
}
