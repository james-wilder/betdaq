package main

import (
	"fmt"
	"log"

	"github.com/james-wilder/betdaq/client"
	"github.com/james-wilder/betdaq/config"
	"github.com/james-wilder/betdaq/model"
)

var configFilename = "./config.json"

func main() {
	conf, err := config.ReadConfig(configFilename)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't load config file" + configFilename)
	}

	client := betdaq.NewClient(conf.Username, conf.Password)

	testGetOddsLadder(client)
	//testGetAccountBalances(client)
	//testGetTopLevelEvents(client)
	//testGetEventSubTreeNoSelections(client, 100004) // Horse Racing
	//testGetMarketInformation(client, 12208599)
	//testGetPrices(client, 12208599)
}

func testGetOddsLadder(c *betdaq.BetdaqClient) {
	getLadderResponse, err := c.GetOddsLadder(model.GetOddsLadder{
		GetOddsLadderRequest: model.GetOddsLadderRequest{
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

func testGetAccountBalances(c *betdaq.BetdaqClient) {
	getAccountBalancesResponse, err := c.GetAccountBalances(model.GetAccountBalances{})
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the account balances")
	}
	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.Currency)
	fmt.Println(getAccountBalancesResponse.GetAccountBalancesResult.AvailableFunds)
}

func testGetTopLevelEvents(c *betdaq.BetdaqClient) {
	getTopLevelEvents, err := c.ListTopLevelEvents(model.ListTopLevelEvents{})
	if err != nil {
		log.Fatal(err)
		panic("Couldn't get the top level events")
	}

	traverseEvents(getTopLevelEvents.ListTopLevelEventsResult.EventClassifiers, "")
}

func traverseEvents(eventClassifiers []model.EventClassifierType, indent string) {
	for _, eventClassifier := range eventClassifiers {
		fmt.Println(indent+"Event", eventClassifier.Id, eventClassifier.Name)
		for _, marketType := range eventClassifier.Markets {
			fmt.Println(indent, marketType.Id, marketType.Name, marketType.Type)
		}
		//fmt.Printf(indent+"Has %d sub types\n", len(eventClassifier.EventClassifiers))
		traverseEvents(eventClassifier.EventClassifiers, indent+"  ")
	}
}

func testGetEventSubTreeNoSelections(c *betdaq.BetdaqClient, id int64) {
	getEventSubTreeNoSelections, err := c.GetEventSubTreeNoSelections(model.GetEventSubTreeNoSelections{
		GetEventSubTreeNoSelectionsRequest: model.GetEventSubTreeNoSelectionsRequest{
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

func testGetMarketInformation(c *betdaq.BetdaqClient, id int64) {
	getMarketInformation, err := c.GetMarketInformation(model.GetMarketInformation{
		GetMarketInformationRequest: model.GetMarketInformationRequest{
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

func testGetPrices(c *betdaq.BetdaqClient, id int64) {
	getPrices, err := c.GetPrices(model.GetPrices{
		GetPricesRequest: model.GetPricesRequest{
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
