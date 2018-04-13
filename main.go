package main

import (
	"fmt"
	"github.com/james-wilder/betdaq/api"
	"github.com/james-wilder/betdaq/config"
	"log"
)

var configFilename = "./config.json"

func main() {
	config, err := config.ReadConfig(configFilename)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't load config file" + configFilename)
	}

	client := api.NewClient(config.Username, config.Password)

	getLadderResponse, err := client.GetOddsLadder(api.PriceFormatDecimal)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the odds ladder")
	}

	for _, price := range getLadderResponse.GetOddsLadderResult.Prices {
		fmt.Println(price.Price, price.Representation)
	}

	//getAccountBalancesResponse, err := client.GetAccountBalances(api.PriceFormatDecimal)
	//if err != nil {
	//	log.Fatal(err)
	//	panic("Couldn't get the account balances")
	//}
	//
	//fmt.Println(getAccountBalancesResponse.Currency)
}
