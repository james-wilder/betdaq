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

	//testGetOddsLadder(client)
	//testGetAccountBalances(client)
	testGetTopLevelEvents(client)
}

func testGetOddsLadder(client *api.Client) {
	getLadderResponse, err := client.GetOddsLadder(1)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the odds ladder")
	}
	for _, price := range getLadderResponse.GetOddsLadderResult.Ladder {
		fmt.Println(price.Price, price.Representation)
	}
}

func testGetAccountBalances(client *api.Client) {
	getAccountBalancesResponse, err := client.GetAccountBalances(1)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the account balances")
	}
	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.Currency)
	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.AvailableFunds)
}

func testGetTopLevelEvents(client *api.Client) {
	getTopLevelEvents, err := client.GetTopLevelEvents()
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the account balances")
	}

	traverseEvents(getTopLevelEvents.ListTopLevelEventsResult.EventClassifiers, "")
}

func traverseEvents(eventClassifiers []api.EventClassifierType, indent string) {
	for _, eventClassifier := range eventClassifiers {
		fmt.Println(indent+"Event", eventClassifier.Id, eventClassifier.Name)
		for _, marketType := range eventClassifier.Markets {
			fmt.Println(indent, marketType.Id, marketType.Name, marketType.Type)
		}
		traverseEvents(eventClassifier.EventClassifiers, indent+"  ")
	}
}
