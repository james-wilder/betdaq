package model

import (
	"fmt"
	"testing"

	"github.com/james-wilder/betdaq/soap"
	"io/ioutil"
)

func TestEncode(t *testing.T) {
	req := &GetPrices{
		GetPricesRequest: GetPricesRequest{
			ThresholdAmount:             "0",
			NumberForPricesRequired:     -1,
			NumberAgainstPricesRequired: -1,
			MarketIds: []int64{
				483492,
			},
		},
	}
	data, err := soap.Encode(&req, "username", "xxx")
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(data))
	if string(data) != expectedSoapRequest {
		t.Log("Not expected SOAP request output")
		t.Fail()
	}
}

func TestDecode(t *testing.T) {
	t.Run("Test GetOddsLadder response", func(t *testing.T) {
		var resp GetOddsLadderResponse

		response, err := ioutil.ReadFile("raw/get-odds-ladder.xml")
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		err = soap.Decode(response, &resp)
		fmt.Println(resp)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if len(resp.GetOddsLadderResult.ReturnStatus) == 0 {
			t.Fail()
			return
		}
		fmt.Println(resp.GetOddsLadderResult.ReturnStatus[0].CallId)
		if resp.GetOddsLadderResult.ReturnStatus[0].CallId != "26091ffa-e9e7-437a-aaf5-6e690bc3e33a" {
			t.Fail()
		}
		fmt.Println("Ladders:", len(resp.GetOddsLadderResult.Ladder))
		if len(resp.GetOddsLadderResult.Ladder) != 495 {
			t.Fail()
		}
		fmt.Println("Ladders[3]:", resp.GetOddsLadderResult.Ladder[3])
		if resp.GetOddsLadderResult.Ladder[3].Price != "1.04" {
			t.Fail()
		}
		if resp.GetOddsLadderResult.Ladder[3].Representation != "1.04" {
			t.Fail()
		}
	})

	t.Run("Test GetPrices response", func(t *testing.T) {
		var resp GetPricesResponse

		response, err := ioutil.ReadFile("raw/get-prices.xml")
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		err = soap.Decode(response, &resp)
		fmt.Println(resp)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if len(resp.GetPricesResult.ReturnStatus) == 0 {
			t.Fail()
			return
		}
		fmt.Println(resp.GetPricesResult.ReturnStatus[0].CallId)
		if resp.GetPricesResult.ReturnStatus[0].CallId != "af26df13-5f2e-43ca-8de5-129893764283" {
			t.Fail()
		}
		fmt.Println("Selections:", len(resp.GetPricesResult.MarketPrices[0].Selections))
		if len(resp.GetPricesResult.MarketPrices[0].Selections) != 8 {
			t.Fail()
		}
		fmt.Println("3rd selection ID:", resp.GetPricesResult.MarketPrices[0].Selections[3].Id)
		if resp.GetPricesResult.MarketPrices[0].Selections[3].Id != 77561824 {
			t.Fail()
		}
	})
}

var expectedSoapRequest = `<?xml version="1.0" encoding="UTF-8"?>
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <Header xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <ExternalApiHeader xmlns="http://www.GlobalBettingExchange.com/ExternalAPI/" version="2" languageCode="en" username="username" password="xxx"></ExternalApiHeader>
  </Header>
  <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <GetPrices xmlns="http://www.GlobalBettingExchange.com/ExternalAPI/">
      <getPricesRequest ThresholdAmount="0" NumberForPricesRequired="-1" NumberAgainstPricesRequired="-1" WantMarketMatchedAmount="false" WantSelectionsMatchedAmounts="false" WantSelectionMatchedDetails="false">
        <MarketIds>483492</MarketIds>
      </getPricesRequest>
    </GetPrices>
</Body>
</Envelope>`
