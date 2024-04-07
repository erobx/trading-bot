package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Testing Steam Web API
	// steamApi()

	m, err := NewMarket()
	if err != nil {
		// Fail to connect to DB
		panic(err)
	}

	svc := NewMarketService(m)
	svc = NewLogService(svc)
	user := NewUser(svc, 100.00)

	app := NewApp(user)
	app.start()
}

func steamApi() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// apiKey := os.Getenv("STEAM_WEB_API_KEY")

	// classIds := [2]string{"613589849", "5710093913"}
	// getAssetClassInfo("2", classIds[:], apiKey)

	// getAssetPrices(apiKey)
}
