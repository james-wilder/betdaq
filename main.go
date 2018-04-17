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

	client := api.NewBetdaqClient(config.Username, config.Password)

	//testGetOddsLadder(client)
	//testGetAccountBalances(client)
	//testGetTopLevelEvents(client)
	//testGetEventSubTreeNoSelections(client, 100004) // Horse Racing
	//testGetMarketInformation(client, 12208599)
	testGetPrices(client, 12208599)
}

func testGetOddsLadder(client *api.BetdaqClient) {
	getLadderResponse, err := client.CallGetOddsLadder(api.GetOddsLadder{
		GetOddsLadderRequest: api.GetOddsLadderRequest{
			PriceFormat: 1,
		},
	})
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the odds ladder")
	}
	for _, price := range getLadderResponse.GetOddsLadderResult.Ladder {
		fmt.Println(price.Price, price.Representation)
	}
}

func testGetAccountBalances(client *api.BetdaqClient) {
	getAccountBalancesResponse, err := client.CallGetAccountBalances(api.GetAccountBalances{})
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the account balances")
	}
	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.Currency)
	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.AvailableFunds)
}

func testGetTopLevelEvents(client *api.BetdaqClient) {
	getTopLevelEvents, err := client.CallListTopLevelEvents(api.ListTopLevelEvents{})
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

func testGetEventSubTreeNoSelections(client *api.BetdaqClient, id int64) {
	getEventSubTreeNoSelections, err := client.CallGetEventSubTreeNoSelections(api.GetEventSubTreeNoSelections{
		GetEventSubTreeNoSelectionsRequest: api.GetEventSubTreeNoSelectionsRequest{
			EventClassifierIds: []int64{
				id,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
		panic("Couldn't do GetEventSubTreeNoSelections")
	}

	traverseEvents(getEventSubTreeNoSelections.GetEventSubTreeNoSelectionsResult.EventClassifiers, "")
}

func testGetMarketInformation(client *api.BetdaqClient, id int64) {
	getMarketInformation, err := client.CallGetMarketInformation(api.GetMarketInformation{
		GetMarketInformationRequest: api.GetMarketInformationRequest{
			MarketIds: []int64{
				id,
			},
		},
	})
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

func testGetPrices(client *api.BetdaqClient, id int64) {
	getPrices, err := client.CallGetPrices(api.GetPrices{
		GetPricesRequest: api.GetPricesRequest{
			MarketIds:                    []int64{id},
			ThresholdAmount:              "0",
			NumberAgainstPricesRequired:  3,
			NumberForPricesRequired:      3,
			WantMarketMatchedAmount:      true,
			WantSelectionMatchedDetails:  true,
			WantSelectionsMatchedAmounts: true,
		},
	})

	if err != nil {
		log.Fatal(err)
		panic("Couldn't do GetMarketInformation")
	}

	for _, market := range getPrices.GetPricesResult.MarketPrices {
		fmt.Println(market.Id, market.Name, market.Type, market.Status, market.StartTime)
		for _, selection := range market.Selections {
			fmt.Println("  ", selection.Id, selection.Name, selection.Status)
			for _, price := range selection.AgainstSidePrices {
				fmt.Println("    Against", price.Price, price.Stake)
			}
			for _, price := range selection.ForSidePrices {
				fmt.Println("    For", price.Price, price.Stake)
			}
		}
	}
}
