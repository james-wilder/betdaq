package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/james-wilder/betdaq/soap"
)

const (
	readOnlyService = "http://api.betdaq.com/v2.0/ReadOnlyService.asmx"
	secureService   = "https://api.betdaq.com/v2.0/Secure/SecureService.asmx"
)

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

	var (
		request = GetOddsLadder{
			GetOddsLadderRequest: GetOddsLadderRequest{
				PriceFormat: format,
			},
		}
		response GetOddsLadderResponse
	)

	err := c.doRequest(request, &response, readOnlyService)
	if err != nil {
		return nil, err
	}

	if response.GetOddsLadderResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetOddsLadderResult.ReturnStatus[0].Code,
			response.GetOddsLadderResult.ReturnStatus[0].Description,
			response.GetOddsLadderResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, err
}

func (c *Client) GetAccountBalances(format int64) (*GetAccountBalancesResponse, error) {
	fmt.Println("GetAccountBalances")

	var (
		request = GetAccountBalances{
			GetAccountBalancesRequest: GetAccountBalancesRequest{},
		}
		response GetAccountBalancesResponse
	)

	err := c.doRequest(request, &response, secureService)
	if err != nil {
		return nil, err
	}

	if response.GetAccountBalancesResult.ReturnStatus[0].Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			response.GetAccountBalancesResult.ReturnStatus[0].Code,
			response.GetAccountBalancesResult.ReturnStatus[0].Description,
			response.GetAccountBalancesResult.ReturnStatus[0].ExtraInformation)
	}

	return &response, nil
}

func (c *Client) doRequest(request, response interface{}, url string) error {
	soapRequest, err := soap.Encode(request, c.Username, c.Password)
	fmt.Println(string(soapRequest))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(soapRequest))

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("HTTP response status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = soap.Decode(body, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
