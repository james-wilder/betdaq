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

	getLadderResponse, err := client.GetOddsLadder(1)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the odds ladder")
	}

	for _, price := range getLadderResponse.GetOddsLadderResult.Ladder {
		fmt.Println(price.Price, price.Representation)
	}

	getAccountBalancesResponse, err := client.GetAccountBalances(1)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the account balances")
	}

	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.Currency)
	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.AvailableFunds)
}
