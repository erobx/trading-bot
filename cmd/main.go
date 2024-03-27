package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"trading-bot/model"

	"github.com/joho/godotenv"
)

var (
	tf2AppId = "440"
	cs2AppId = "730"
)

func main() {
	// Testing Steam Web API
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("STEAM_WEB_API_KEY")

	// classIds := [2]string{"195151", "16891096"}
	// getAssetClassInfo("2", classIds[:], apiKey)

	getAssetPrices(apiKey)
}

func getAssetClassInfo(class_count string, classIds []string, key string) {
	req := model.NewRequest().
		WithGenInterface("ISteamEconomy").
		WithMethod("GetAssetClassInfo").
		WithApiKey(key).
		WithAppId(tf2AppId)

	reqUrl := req.CombineUrl() + "class_count=" + class_count + "&"
	for i, id := range classIds {
		reqUrl = reqUrl + "classid" + strconv.Itoa(i) + "=" + id
		if i != len(classIds)-1 {
			reqUrl = reqUrl + "&"
		}
	}

	placeHolder(reqUrl)
}

func getAssetPrices(key string) {
	req := model.NewRequest().
		WithGenInterface("ISteamEconomy").
		WithMethod("GetAssetPrices").
		WithApiKey(key).
		WithAppId(cs2AppId)

	reqUrl := req.CombineUrl() + "languange=en&currency=USD"

	placeHolder(reqUrl)
}

func placeHolder(reqUrl string) {
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	log.Printf(sb)
}

type Reponse interface {
	Parse() error
}

type AssetClassInfo struct {
}

type AssetPrices struct {
	result Result `json:"result"`
}

type Result struct {
}

func (assetCI AssetClassInfo) Parse() error {

	return nil
}

func (ap AssetPrices) Parse() error {
	return nil
}
