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
	//testGetTopLevelEvents(client)
	//testGetEventSubTreeNoSelections(client, 100004) // Horse Racing
	testGetMarketInformation(client, 12196309)
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
		panic("Couldn't get the top level events")
	}

	traverseEvents(getTopLevelEvents.ListTopLevelEventsResult.EventClassifiers, "")
}

func traverseEvents(eventClassifiers []api.EventClassifierType, indent string) {
	for _, eventClassifier := range eventClassifiers {
		fmt.Println(indent+"Event", eventClassifier.Id, eventClassifier.Name)
		for _, marketType := range eventClassifier.Markets {
			fmt.Println(indent, marketType.Id, marketType.Name, marketType.Type)
		}
		//fmt.Printf(indent+"Has %d sub types\n", len(eventClassifier.EventClassifiers))
		traverseEvents(eventClassifier.EventClassifiers, indent+"  ")
	}
}

func testGetEventSubTreeNoSelections(client *api.Client, id int64) {
	getEventSubTreeNoSelections, err := client.GetEventSubTreeNoSelections(id)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't do GetEventSubTreeNoSelections")
	}

	traverseEvents(getEventSubTreeNoSelections.GetEventSubTreeNoSelectionsResult.EventClassifiers, "")
}

func testGetMarketInformation(client *api.Client, id int64) {
	getMarketInformation, err := client.GetMarketInformation(id)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't do GetMarketInformation")
	}

	for _, market := range getMarketInformation.GetMarketInformationResult.Markets {
		fmt.Println(market.Id, market.Name, market.Type, market.Status, market.StartTime)
		for _, selection := range market.Selections {
			fmt.Println("  ", selection.Id, selection.Name, selection.Status)
		}
	}
}
