package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/james-wilder/betdaq/soap"
	"io/ioutil"
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

func (c *Client) GetOddsLadder(format PriceFormat) (*GetOddsLadderResponse, error) {
	fmt.Println("GetOddsLadder")

	request := GetOddsLadder{
		GetOddsLadderRequest: GetOddsLadderRequest{
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

	if content.GetOddsLadderResult.ReturnStatus.Code != 0 {
		return nil, fmt.Errorf("API returned code %d (description:\"%s\", extra information:\"%s\")",
			content.GetOddsLadderResult.ReturnStatus.Code,
			content.GetOddsLadderResult.ReturnStatus.Description,
			content.GetOddsLadderResult.ReturnStatus.ExtraInformation)
	}

	return &content, nil
}
