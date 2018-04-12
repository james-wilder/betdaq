package api

/*
  GetPrices
*/

// GetPrices request root object
type GetPrices struct {
	XMLName          struct{}         `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetPrices"`
	GetPricesRequest GetPricesRequest `xml:"getPricesRequest,allowempty"`
}

// GetPricesRequest contains the content for GetPrices API request
type GetPricesRequest struct {
	ThresholdAmount             int32 `xml:"ThresholdAmount,attr"`
	NumberForPricesRequired     int32 `xml:"NumberForPricesRequired,attr"`
	NumberAgainstPricesRequired int32 `xml:"NumberAgainstPricesRequired,attr"`
	MarketIds                   string
}

// GetPricesResponse response root object
type GetPricesResponse struct {
	XMLName         struct{} `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetPricesResponse"`
	GetPricesResult GetPricesResult
}

// GetPricesResult contains the content for GetPricesResponse API response
type GetPricesResult struct {
	ReturnStatus ReturnStatus
}

// ReturnStatus contains standard API response
type ReturnStatus struct {
	Code             int    `xml:",attr"`
	Description      string `xml:",attr"`
	CallId           string `xml:",attr"`
	ExtraInformation string `xml:",attr"`
}

/*
GetOddsLadder
*/

type GetOddsLadder struct {
	XMLName              struct{}             `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadder"`
	GetOddsLadderRequest GetOddsLadderRequest `xml:"getOddsLadderRequest,allowempty"`
}

type GetOddsLadderRequest struct {
	PriceFormat PriceFormat `xml:"PriceFormat,attr"`
}

type GetOddsLadderResponse struct {
	XMLName             struct{} `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadderResponse"`
	GetOddsLadderResult GetOddsLadderResult
}

type GetOddsLadderResult struct {
	ReturnStatus ReturnStatus
	Prices       []Ladder `xml:"Ladder"`
}

type Ladder struct {
	Price          string `xml:"price,attr"`
	Representation string `xml:"representation,attr"`
}

type PriceFormat int32

const (
	PriceFormatDecimal    PriceFormat = 1
	PriceFormatFractional PriceFormat = 2
	PriceFormatAmerican   PriceFormat = 3
)
