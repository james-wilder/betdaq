package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/james-wilder/betdaq/soap"
)

const readOnlyService = "http://api.betdaq.com/v2.0/ReadOnlyService.asmx"
const secureService = "https://api.betdaq.com/v2.0/Secure/SecureService.asmx"
const port = 443 // HTTPS

type Client struct {
	Username string
	Password string
}

func NewClient(username string, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

func (c *Client) GetOddsLadder(format int64) (*GetOddsLadderResponse, error) {
	fmt.Println("GetOddsLadder")

	request := GetOddsLadder{
		getOddsLadderRequest: GetOddsLadderRequest{
			PriceFormat: format,
		},
	}

	soapRequest, err := soap.Encode(request, c.Username, c.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", readOnlyService, bytes.NewBuffer(soapRequest))

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("HTTP response status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var content GetOddsLadderResponse
	err = soap.Decode(body, &content)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if content.GetOddsLadderResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			content.GetOddsLadderResult.ReturnStatus[0].Code,
			content.GetOddsLadderResult.ReturnStatus[0].Description,
			content.GetOddsLadderResult.ReturnStatus[0].ExtraInformation)
	}

	return &content, nil
}

func (c *Client) GetAccountBalances(format int64) (*GetAccountBalancesResponse, error) {
	fmt.Println("GetAccountBalances")

	request := GetAccountBalances{
		getAccountBalancesRequest: GetAccountBalancesRequest{},
	}

	soapRequest, err := soap.Encode(request, c.Username, c.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(soapRequest))

	req, err := http.NewRequest("POST", secureService, bytes.NewBuffer(soapRequest))

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("HTTP response status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var content GetAccountBalancesResponse
	err = soap.Decode(body, &content)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if content.GetAccountBalancesResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			content.GetAccountBalancesResult.ReturnStatus[0].Code,
			content.GetAccountBalancesResult.ReturnStatus[0].Description,
			content.GetAccountBalancesResult.ReturnStatus[0].ExtraInformation)
	}

	return &content, nil
}
