package api

type PriceFormat int64

const (
	PriceFormatDecimal    PriceFormat = 1
	PriceFormatFractional PriceFormat = 2
	PriceFormatAmerican   PriceFormat = 3
)

/*
  GetPrices
*/

// GetPrices request root object
type GetPrices struct {
	XMLName          struct{}         `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetPrices"`
	GetPricesRequest GetPricesRequest `xml:"getPricesRequest,allowempty"`
}

// GetPricesResult contains the content for GetPricesResponse API response
type GetPricesResult struct {
	ReturnStatus ReturnStatus
}

type GetOddsLadder struct {
	XMLName              struct{}             `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadder"`
	GetOddsLadderRequest GetOddsLadderRequest `xml:"getOddsLadderRequest,allowempty"`
}

type GetOddsLadderResult struct {
	ReturnStatus ReturnStatus
	Prices       []Ladder `xml:"Ladder"`
}

type Ladder struct {
	Price          string `xml:"price,attr"`
	Representation string `xml:"representation,attr"`
}

type GetAccountBalances struct {
	XMLName                   struct{} `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetAccountBalancesRequest"`
	GetAccountBalancesRequest GetAccountBalancesRequest
}
