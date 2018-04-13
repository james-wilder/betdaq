package wsdl

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type PolarityEnum string

const (
	PolarityEnumFor PolarityEnum = "For"

	PolarityEnumAgainst PolarityEnum = "Against"
)

type ListTopLevelEvents struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListTopLevelEvents"`

	ListTopLevelEventsRequest *ListTopLevelEventsRequest `xml:"listTopLevelEventsRequest,omitempty"`
}

type ListTopLevelEventsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListTopLevelEventsResponse"`

	ListTopLevelEventsResult *ListTopLevelEventsResponse `xml:"ListTopLevelEventsResult,omitempty"`
}

type GetPrices struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetPrices"`

	GetPricesRequest *GetPricesRequest `xml:"getPricesRequest,omitempty"`
}

type GetPricesResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetPricesResponse"`

	GetPricesResult *GetPricesResponse `xml:"GetPricesResult,omitempty"`
}

type ListMarketWithdrawalHistory struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListMarketWithdrawalHistory"`

	ListMarketWithdrawalHistoryRequest *ListMarketWithdrawalHistoryRequest `xml:"listMarketWithdrawalHistoryRequest,omitempty"`
}

type ListMarketWithdrawalHistoryResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListMarketWithdrawalHistoryResponse"`

	ListMarketWithdrawalHistoryResult *ListMarketWithdrawalHistoryResponse `xml:"ListMarketWithdrawalHistoryResult,omitempty"`
}

type GetEventSubTreeWithSelections struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeWithSelections"`

	GetEventSubTreeWithSelectionsRequest *GetEventSubTreeWithSelectionsRequest `xml:"getEventSubTreeWithSelectionsRequest,omitempty"`
}

type GetEventSubTreeWithSelectionsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeWithSelectionsResponse"`

	GetEventSubTreeWithSelectionsResult *GetEventSubTreeWithSelectionsResponse `xml:"GetEventSubTreeWithSelectionsResult,omitempty"`
}

type ListOrdersChangedSince struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListOrdersChangedSince"`

	ListOrdersChangedSinceRequest *ListOrdersChangedSinceRequest `xml:"listOrdersChangedSinceRequest,omitempty"`
}

type ListOrdersChangedSinceResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListOrdersChangedSinceResponse"`

	ListOrdersChangedSinceResult *ListOrdersChangedSinceResponse `xml:"ListOrdersChangedSinceResult,omitempty"`
}

type GetEventSubTreeNoSelections struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeNoSelections"`

	GetEventSubTreeNoSelectionsRequest *GetEventSubTreeNoSelectionsRequest `xml:"getEventSubTreeNoSelectionsRequest,omitempty"`
}

type GetEventSubTreeNoSelectionsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeNoSelectionsResponse"`

	GetEventSubTreeNoSelectionsResult *GetEventSubTreeNoSelectionsResponse `xml:"GetEventSubTreeNoSelectionsResult,omitempty"`
}

type GetMarketInformation struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetMarketInformation"`

	GetMarketInformationRequest *GetMarketInformationRequest `xml:"getMarketInformationRequest,omitempty"`
}

type GetMarketInformationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetMarketInformationResponse"`

	GetMarketInformationResult *GetMarketInformationResponse `xml:"GetMarketInformationResult,omitempty"`
}

type PlaceOrdersNoReceipt struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersNoReceipt"`

	PlaceOrdersNoReceiptRequest *PlaceOrdersNoReceiptRequest `xml:"placeOrdersNoReceiptRequest,omitempty"`
}

type PlaceOrdersNoReceiptResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersNoReceiptResponse"`

	PlaceOrdersNoReceiptResult *PlaceOrdersNoReceiptResponse `xml:"PlaceOrdersNoReceiptResult,omitempty"`
}

type PlaceOrdersWithReceipt struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersWithReceipt"`

	Orders *PlaceOrdersWithReceiptRequest `xml:"orders,omitempty"`
}

type PlaceOrdersWithReceiptResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersWithReceiptResponse"`

	PlaceOrdersWithReceiptResult *PlaceOrdersWithReceiptResponse `xml:"PlaceOrdersWithReceiptResult,omitempty"`
}

type CancelOrders struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelOrders"`

	CancelOrdersRequest *CancelOrdersRequest `xml:"cancelOrdersRequest,omitempty"`
}

type CancelOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelOrdersResponse"`

	CancelOrdersResult *CancelOrdersResponse `xml:"CancelOrdersResult,omitempty"`
}

type CancelAllOrders struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrders"`

	CancelAllOrdersRequest *CancelAllOrdersRequest `xml:"cancelAllOrdersRequest,omitempty"`
}

type CancelAllOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersResponse"`

	CancelAllOrdersResult *CancelAllOrdersResponse `xml:"CancelAllOrdersResult,omitempty"`
}

type CancelAllOrdersOnMarket struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersOnMarket"`

	CancelAllOrdersOnMarketRequest *CancelAllOrdersOnMarketRequest `xml:"cancelAllOrdersOnMarketRequest,omitempty"`
}

type CancelAllOrdersOnMarketResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersOnMarketResponse"`

	CancelAllOrdersOnMarketResult *CancelAllOrdersOnMarketResponse `xml:"CancelAllOrdersOnMarketResult,omitempty"`
}

type GetAccountBalances struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetAccountBalances"`

	GetAccountBalancesRequest *GetAccountBalancesRequest `xml:"getAccountBalancesRequest,omitempty"`
}

type GetAccountBalancesResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetAccountBalancesResponse"`

	GetAccountBalancesResult *GetAccountBalancesResponse `xml:"GetAccountBalancesResult,omitempty"`
}

type ListAccountPostings struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostings"`

	ListAccountPostingsRequest *ListAccountPostingsRequest `xml:"listAccountPostingsRequest,omitempty"`
}

type ListAccountPostingsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsResponse"`

	ListAccountPostingsResult *ListAccountPostingsResponse `xml:"ListAccountPostingsResult,omitempty"`
}

type ListAccountPostingsById struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsById"`

	ListAccountPostingsByIdRequest *ListAccountPostingsByIdRequest `xml:"listAccountPostingsByIdRequest,omitempty"`
}

type ListAccountPostingsByIdResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsByIdResponse"`

	ListAccountPostingsByIdResult *ListAccountPostingsByIdResponse `xml:"ListAccountPostingsByIdResult,omitempty"`
}

type UpdateOrdersNoReceipt struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UpdateOrdersNoReceipt"`

	UpdateOrdersNoReceiptRequest *UpdateOrdersNoReceiptRequest `xml:"updateOrdersNoReceiptRequest,omitempty"`
}

type UpdateOrdersNoReceiptResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UpdateOrdersNoReceiptResponse"`

	UpdateOrdersNoReceiptResult *UpdateOrdersNoReceiptResponse `xml:"UpdateOrdersNoReceiptResult,omitempty"`
}

type GetOrderDetails struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOrderDetails"`

	GetOrderDetailsRequest *GetOrderDetailsRequest `xml:"getOrderDetailsRequest,omitempty"`
}

type GetOrderDetailsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOrderDetailsResponse"`

	GetOrderDetailsResult *GetOrderDetailsResponse `xml:"GetOrderDetailsResult,omitempty"`
}

type ChangePassword struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangePassword"`

	ChangePasswordRequest *ChangePasswordRequest `xml:"changePasswordRequest,omitempty"`
}

type ChangePasswordResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangePasswordResponse"`

	ChangePasswordResult *ChangePasswordResponse `xml:"ChangePasswordResult,omitempty"`
}

type ListSelectionsChangedSince struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionsChangedSince"`

	ListSelectionsChangedSinceRequest *ListSelectionsChangedSinceRequest `xml:"listSelectionsChangedSinceRequest,omitempty"`
}

type ListSelectionsChangedSinceResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionsChangedSinceResponse"`

	ListSelectionsChangedSinceResult *ListSelectionsChangedSinceResponse `xml:"ListSelectionsChangedSinceResult,omitempty"`
}

type ListBootstrapOrders struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBootstrapOrders"`

	ListBootstrapOrdersRequest *ListBootstrapOrdersRequest `xml:"listBootstrapOrdersRequest,omitempty"`
}

type ListBootstrapOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBootstrapOrdersResponse"`

	ListBootstrapOrdersResult *ListBootstrapOrdersResponse `xml:"ListBootstrapOrdersResult,omitempty"`
}

type SuspendFromTrading struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendFromTrading"`

	SuspendFromTradingRequest *SuspendFromTradingRequest `xml:"suspendFromTradingRequest,omitempty"`
}

type SuspendFromTradingResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendFromTradingResponse"`

	SuspendFromTradingResult *SuspendFromTradingResponse `xml:"SuspendFromTradingResult,omitempty"`
}

type UnsuspendFromTrading struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendFromTrading"`

	UnsuspendFromTradingRequest *UnsuspendFromTradingRequest `xml:"unsuspendFromTradingRequest,omitempty"`
}

type UnsuspendFromTradingResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendFromTradingResponse"`

	UnsuspendFromTradingResult *UnsuspendFromTradingResponse `xml:"UnsuspendFromTradingResult,omitempty"`
}

type SuspendOrders struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendOrders"`

	SuspendOrdersRequest *SuspendOrdersRequest `xml:"suspendOrdersRequest,omitempty"`
}

type SuspendOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendOrdersResponse"`

	SuspendOrdersResult *SuspendOrdersResponse `xml:"SuspendOrdersResult,omitempty"`
}

type SuspendAllOrdersOnMarket struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrdersOnMarket"`

	SuspendAllOrdersOnMarket *SuspendAllOrdersOnMarketRequest `xml:"suspendAllOrdersOnMarket,omitempty"`
}

type SuspendAllOrdersOnMarketResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrdersOnMarketResponse"`

	SuspendAllOrdersOnMarketResult *SuspendAllOrdersOnMarketResponse `xml:"SuspendAllOrdersOnMarketResult,omitempty"`
}

type UnsuspendOrders struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendOrders"`

	UnsuspendOrdersRequest *UnsuspendOrdersRequest `xml:"unsuspendOrdersRequest,omitempty"`
}

type UnsuspendOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendOrdersResponse"`

	UnsuspendOrdersResult *UnsuspendOrdersResponse `xml:"UnsuspendOrdersResult,omitempty"`
}

type SuspendAllOrders struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrders"`

	SuspendAllOrdersRequest *SuspendAllOrdersRequest `xml:"suspendAllOrdersRequest,omitempty"`
}

type SuspendAllOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrdersResponse"`

	SuspendAllOrdersResult *SuspendAllOrdersResponse `xml:"SuspendAllOrdersResult,omitempty"`
}

type ListBlacklistInformation struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBlacklistInformation"`

	ListBlacklistInformationRequest *ListBlacklistInformationRequest `xml:"listBlacklistInformationRequest,omitempty"`
}

type ListBlacklistInformationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBlacklistInformationResponse"`

	ListBlacklistInformationResult *ListBlacklistInformationResponse `xml:"ListBlacklistInformationResult,omitempty"`
}

type RegisterHeartbeat struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ RegisterHeartbeat"`

	RegisterHeartbeatRequest *RegisterHeartbeatRequest `xml:"registerHeartbeatRequest,omitempty"`
}

type RegisterHeartbeatResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ RegisterHeartbeatResponse"`

	RegisterHeartbeatResult *RegisterHeartbeatResponse `xml:"RegisterHeartbeatResult,omitempty"`
}

type ChangeHeartbeatRegistration struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangeHeartbeatRegistration"`

	ChangeHeartbeatRegistrationRequest *ChangeHeartbeatRegistrationRequest `xml:"changeHeartbeatRegistrationRequest,omitempty"`
}

type ChangeHeartbeatRegistrationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangeHeartbeatRegistrationResponse"`

	ChangeHeartbeatRegistrationResult *ChangeHeartbeatRegistrationResponse `xml:"ChangeHeartbeatRegistrationResult,omitempty"`
}

type DeregisterHeartbeat struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ DeregisterHeartbeat"`

	DeregisterHeartbeatRequest *DeregisterHeartbeatRequest `xml:"deregisterHeartbeatRequest,omitempty"`
}

type DeregisterHeartbeatResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ DeregisterHeartbeatResponse"`

	DeregisterHeartbeatResult *DeregisterHeartbeatResponse `xml:"DeregisterHeartbeatResult,omitempty"`
}

type Pulse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ Pulse"`

	PulseRequest *PulseRequest `xml:"pulseRequest,omitempty"`
}

type PulseResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PulseResponse"`

	PulseResult *PulseResponse `xml:"PulseResult,omitempty"`
}

type GetOddsLadder struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadder"`

	GetOddsLadderRequest *GetOddsLadderRequest `xml:"getOddsLadderRequest,omitempty"`
}

type GetOddsLadderResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadderResponse"`

	GetOddsLadderResult *GetOddsLadderResponse `xml:"GetOddsLadderResult,omitempty"`
}

type GetCurrentSelectionSequenceNumber struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetCurrentSelectionSequenceNumber"`

	GetCurrentSelectionSequenceNumberRequest *GetCurrentSelectionSequenceNumberRequest `xml:"getCurrentSelectionSequenceNumberRequest,omitempty"`
}

type GetCurrentSelectionSequenceNumberResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetCurrentSelectionSequenceNumberResponse"`

	GetCurrentSelectionSequenceNumberResult *GetCurrentSelectionSequenceNumberResponse `xml:"GetCurrentSelectionSequenceNumberResult,omitempty"`
}

type ListSelectionTrades struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionTrades"`

	ListSelectionTradesRequest *ListSelectionTradesRequest `xml:"listSelectionTradesRequest,omitempty"`
}

type ListSelectionTradesResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionTradesResponse"`

	ListSelectionTradesResult *ListSelectionTradesResponse `xml:"ListSelectionTradesResult,omitempty"`
}

type GetSPEnabledMarketsInformation struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetSPEnabledMarketsInformation"`

	GetSPEnabledMarketsInformationRequest *GetSPEnabledMarketsInformationRequest `xml:"GetSPEnabledMarketsInformationRequest,omitempty"`
}

type GetSPEnabledMarketsInformationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetSPEnabledMarketsInformationResponse"`

	GetSPEnabledMarketsInformationResult *GetSPEnabledMarketsInformationResponse `xml:"GetSPEnabledMarketsInformationResult,omitempty"`
}

type ReturnStatus struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ReturnStatus"`

	Code int32 `xml:"Code,attr,omitempty"`

	Description string `xml:"Description,attr,omitempty"`

	CallId string `xml:"CallId,attr,omitempty"`

	ExtraInformation string `xml:"ExtraInformation,attr,omitempty"`
}

type GetAccountBalancesResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetAccountBalancesResponse"`

	*BaseResponse

	Currency string `xml:"Currency,attr,omitempty"`

	Balance float64 `xml:"Balance,attr,omitempty"`

	Exposure float64 `xml:"Exposure,attr,omitempty"`

	AvailableFunds float64 `xml:"AvailableFunds,attr,omitempty"`

	Credit float64 `xml:"Credit,attr,omitempty"`
}

type ListSelectionsChangedSinceRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionsChangedSinceRequest"`

	SelectionSequenceNumber int64 `xml:"SelectionSequenceNumber,attr,omitempty"`
}

type ListSelectionsChangedSinceResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionsChangedSinceResponse"`

	*BaseResponse

	Selections *ListSelectionsChangedSinceResponseItem `xml:"Selections,omitempty"`
}

type ListSelectionsChangedSinceResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionsChangedSinceResponseItem"`

	SettlementInformation *SettlementInformationType `xml:"SettlementInformation,omitempty"`

	Id int64 `xml:"Id,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	DisplayOrder int32 `xml:"DisplayOrder,attr,omitempty"`

	IsHidden bool `xml:"IsHidden,attr,omitempty"`

	Status int16 `xml:"Status,attr,omitempty"`

	ResetCount int16 `xml:"ResetCount,attr,omitempty"`

	WithdrawalFactor float64 `xml:"WithdrawalFactor,attr,omitempty"`

	MarketId int64 `xml:"MarketId,attr,omitempty"`

	SelectionSequenceNumber int64 `xml:"SelectionSequenceNumber,attr,omitempty"`

	CancelOrdersTime time.Time `xml:"CancelOrdersTime,attr,omitempty"`
}

type SettlementInformationType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SettlementInformationType"`

	SettledTime time.Time `xml:"SettledTime,attr,omitempty"`

	VoidPercentage float64 `xml:"VoidPercentage,attr,omitempty"`

	LeftSideFactor float64 `xml:"LeftSideFactor,attr,omitempty"`

	RightSideFactor float64 `xml:"RightSideFactor,attr,omitempty"`

	SettlementResultString string `xml:"SettlementResultString,attr,omitempty"`
}

type ListBootstrapOrdersRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBootstrapOrdersRequest"`

	SequenceNumber                      int64 `xml:"SequenceNumber,omitempty"`
	WantSettledOrdersOnUnsettledMarkets bool  `xml:"wantSettledOrdersOnUnsettledMarkets,omitempty"`
}

type ListBootstrapOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBootstrapOrdersResponse"`

	*BaseResponse

	Orders struct {
		Order []*Order `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`

	MaximumSequenceNumber int64 `xml:"MaximumSequenceNumber,attr,omitempty"`
}

type Order struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ Order"`

	OrderCommissionInformation *OrderCommissionInformationType `xml:"OrderCommissionInformation,omitempty"`

	Id int64 `xml:"Id,attr,omitempty"`

	MarketId int64 `xml:"MarketId,attr,omitempty"`

	SelectionId int64 `xml:"SelectionId,attr,omitempty"`

	SequenceNumber int64 `xml:"SequenceNumber,attr,omitempty"`

	IssuedAt time.Time `xml:"IssuedAt,attr,omitempty"`

	Polarity byte `xml:"Polarity,attr,omitempty"`

	UnmatchedStake float64 `xml:"UnmatchedStake,attr,omitempty"`

	RequestedPrice float64 `xml:"RequestedPrice,attr,omitempty"`

	MatchedPrice float64 `xml:"MatchedPrice,attr,omitempty"`

	MatchedStake float64 `xml:"MatchedStake,attr,omitempty"`

	TotalForSideMakeStake float64 `xml:"TotalForSideMakeStake,attr,omitempty"`

	TotalForSideTakeStake float64 `xml:"TotalForSideTakeStake,attr,omitempty"`

	MatchedAgainstStake float64 `xml:"MatchedAgainstStake,attr,omitempty"`

	Status byte `xml:"Status,attr,omitempty"`

	RestrictOrderToBroker bool `xml:"RestrictOrderToBroker,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`

	CancelOnInRunning bool `xml:"CancelOnInRunning,attr,omitempty"`

	CancelIfSelectionReset bool `xml:"CancelIfSelectionReset,attr,omitempty"`

	IsCurrentlyInRunning bool `xml:"IsCurrentlyInRunning,attr,omitempty"`

	PunterCommissionBasis byte `xml:"PunterCommissionBasis,attr,omitempty"`

	MakeCommissionRate float64 `xml:"MakeCommissionRate,attr,omitempty"`

	TakeCommissionRate float64 `xml:"TakeCommissionRate,attr,omitempty"`

	ExpectedSelectionResetCount byte `xml:"ExpectedSelectionResetCount,attr,omitempty"`

	ExpectedWithdrawalSequenceNumber byte `xml:"ExpectedWithdrawalSequenceNumber,attr,omitempty"`

	OrderFillType byte `xml:"OrderFillType,attr,omitempty"`

	FillOrKillThreshold float64 `xml:"FillOrKillThreshold,attr,omitempty"`
}

type OrderCommissionInformationType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ OrderCommissionInformationType"`

	GrossSettlementAmount float64 `xml:"GrossSettlementAmount,attr,omitempty"`

	OrderCommission float64 `xml:"OrderCommission,attr,omitempty"`
}

type ArrayOfOrder struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ArrayOfOrder"`

	Order []*Order `xml:"Order,omitempty"`
}

type SimpleOrderRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SimpleOrderRequest"`

	SelectionId int64 `xml:"SelectionId,attr,omitempty"`

	Stake float64 `xml:"Stake,attr,omitempty"`

	Price float64 `xml:"Price,attr,omitempty"`

	Polarity byte `xml:"Polarity,attr,omitempty"`

	// This value must be set to the SelectionResetCount value that is in the Selection object that is returned by GetMarketInformation and GetPrices.  The purpose of this is to ensure that you are aware of the state of the market before placing a bet.  If the ExpectedSelectionResetCount that you submit to the server is not the same as the SelectionResetCount that is on the server, then your bet will NOT be accepted, and you will get an error.
	ExpectedSelectionResetCount byte `xml:"ExpectedSelectionResetCount,attr,omitempty"`

	// This value should be set to the withdrawalSequenceNumber value that is in the Market object that is returned by GetMarketInformation and GetPrices.  The purpose of this is to ensure that you are aware of the state of the market before placing a bet.  If the expectedWithdrawalSequenceNumbert that you submit to the server is not the same as the withdrawalSequenceNumber that is on the server, then your bet WILL be accepted, but it will be repriced to allow for the fact that there are more or less selections available on the market than you believed.
	ExpectedWithdrawalSequenceNumber byte `xml:"ExpectedWithdrawalSequenceNumber,attr,omitempty"`

	// The effect of this option when set to true is to cancel any unmatched orders when the market changes to an in-running market.  This option only applies while the market is NOT in-running.  When the market turns in-running, this option will have no effect.
	CancelOnInRunning bool `xml:"CancelOnInRunning,attr,omitempty"`

	// The effect of this option is to cancel any unmatched bets if the selection is reset.  This can occur when the Market is reset (eg a goal is scored in an in-running market).  This defaults to true - unmatched bets will be cancelled if an event occurs in the market such that Betdaq forces the market to be reset
	CancelIfSelectionReset bool `xml:"CancelIfSelectionReset,attr,omitempty"`

	// An expires at value set in the past will cause the bet to be cancelled - although  the bet status will not immediately be set to Cancelled (this will occur in several moments on the exchange), the bet will not be available for matching.
	ExpiresAt time.Time `xml:"ExpiresAt,attr,omitempty"`

	// Reprice(1), Cancel(2), DontReprice(3)
	WithdrawalRepriceOption byte `xml:"WithdrawalRepriceOption,attr,omitempty"`

	KillType byte `xml:"KillType,attr,omitempty"`

	// This field is required only if killType is set to FillOrKill or FillOrKillDontCancel.
	FillOrKillThreshold float64 `xml:"FillOrKillThreshold,attr,omitempty"`

	// Deprecated - This field is required only if killType is set to FillOrKillDontCancel.
	RestrictOrderToBroker bool `xml:"RestrictOrderToBroker,attr,omitempty"`

	// For internal use only
	ChannelTypeInfo string `xml:"ChannelTypeInfo,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`
}

type Offer struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ Offer"`

	// The amount available for this price.
	Stake float64 `xml:"Stake,omitempty"`

	Price struct {
	} `xml:"Price,omitempty"`
}

type BaseResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ BaseResponse"`

	ReturnStatus *ReturnStatus `xml:"ReturnStatus,omitempty"`
	Timestamp    time.Time     `xml:"Timestamp,omitempty"`
}

type ListOrdersChangedSinceRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListOrdersChangedSinceRequest"`

	SequenceNumber int64 `xml:"SequenceNumber,omitempty"`
}

type ListOrdersChangedSinceResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListOrdersChangedSinceResponse"`

	*BaseResponse

	Orders struct {
		Order []*Order `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type PlaceOrdersNoReceiptRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersNoReceiptRequest"`

	Orders struct {
		Order []*SimpleOrderRequest `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`

	WantAllOrNothingBehaviour bool `xml:"WantAllOrNothingBehaviour,omitempty"`
}

type PlaceOrdersNoReceiptResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersNoReceiptResponseItem"`

	OrderHandle int64 `xml:"OrderHandle,attr,omitempty"`

	ReturnCode int32 `xml:"ReturnCode,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`
}

type PlaceOrdersNoReceiptResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersNoReceiptResponse"`

	*BaseResponse

	OrderHandles struct {
		OrderHandle []int64 `xml:"OrderHandle,omitempty"`
	} `xml:"OrderHandles,omitempty"`

	Orders struct {
		Order []*PlaceOrdersNoReceiptResponseItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type PlaceOrdersWithReceiptRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersWithReceiptRequest"`

	Orders struct {
		Order []*PlaceOrdersWithReceiptRequestItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type PlaceOrdersWithReceiptResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersWithReceiptResponse"`

	*BaseResponse

	Orders struct {
		Order []*PlaceOrdersWithReceiptResponseItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type CancelOrdersResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelOrdersResponseItem"`

	OrderHandle int64 `xml:"OrderHandle,attr,omitempty"`

	CancelledForSideStake float64 `xml:"cancelledForSideStake,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`
}

type CancelOrdersRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelOrdersRequest"`
}

type CancelOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelOrdersResponse"`

	*BaseResponse

	CancelledOrdersHandles struct {
		OrderHandle []int64 `xml:"OrderHandle,omitempty"`
	} `xml:"CancelledOrdersHandles,omitempty"`

	Orders struct {
		Order []*CancelOrdersResponseItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type CancelAllOrdersRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersRequest"`
}

type CancelAllOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersResponse"`

	*BaseResponse

	CancelledOrdersHandles struct {
		OrderHandle []int64 `xml:"OrderHandle,omitempty"`
	} `xml:"CancelledOrdersHandles,omitempty"`

	Orders struct {
		Order []*CancelAllOrdersResponseItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type PlaceOrdersWithReceiptResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersWithReceiptResponseItem"`

	Status byte `xml:"Status,attr,omitempty"`

	MatchedAgainstStake float64 `xml:"MatchedAgainstStake,attr,omitempty"`

	MatchedStake float64 `xml:"MatchedStake,attr,omitempty"`

	MatchedPrice float64 `xml:"MatchedPrice,attr,omitempty"`

	UnmatchedStake float64 `xml:"UnmatchedStake,attr,omitempty"`

	Polarity byte `xml:"Polarity,attr,omitempty"`

	IssuedAt time.Time `xml:"IssuedAt,attr,omitempty"`

	SequenceNumber int64 `xml:"SequenceNumber,attr,omitempty"`

	// Deprecated attribute.
	SelectionId int64 `xml:"SelectionId,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`

	OrderHandle int64 `xml:"OrderHandle,attr,omitempty"`
}

type PlaceOrdersWithReceiptRequestItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PlaceOrdersWithReceiptRequestItem"`

	SelectionId int64 `xml:"SelectionId,attr,omitempty"`

	Stake float64 `xml:"Stake,attr,omitempty"`

	Price float64 `xml:"Price,attr,omitempty"`

	Polarity byte `xml:"Polarity,attr,omitempty"`

	// This value must be set to the selectionResetCount value that is in the Selection object that is returned by GetMarketInformation and GetPrices.  The purpose of this is to ensure that you are aware of the state of the market before placing a bet.  If the expectedSelectionResetCount that you submit to the server is not the same as the selectionResetCount that is on the server, then your bet will NOT be accepted, and you will get a 300 error.
	ExpectedSelectionResetCount byte `xml:"ExpectedSelectionResetCount,attr,omitempty"`

	// This value should be set to the withdrawalSequenceNumber value that is in the Market object that is returned by GetMarketInformation and GetPrices.  The purpose of this is to ensure that you are aware of the state of the market before placing a bet.  If the expectedWithdrawalSequenceNumbert that you submit to the server is not the same as the withdrawalSequenceNumber that is on the server, then your bet WILL be accepted, but it will be repriced to allow for the fact that there are more or less selections available on the market than you believed.
	ExpectedWithdrawalSequenceNumber byte `xml:"ExpectedWithdrawalSequenceNumber,attr,omitempty"`

	// FillAndKill=2, FillOrKill=3, FillOrKillDontCancel=4, SPIfUnmatched=5
	KillType byte `xml:"KillType,attr,omitempty"`

	// This field is required only if killType is set to FillOrKill or FillOrKillDontCancel.
	FillOrKillThreshold float64 `xml:"FillOrKillThreshold,attr,omitempty"`

	//
	// The effect of this option when set to true is to cancel any unmatched bets when the market changes to an in-running market.  This option only applies while the market is NOT in-running.  When the market turns in-running, this option will have no effect.
	// This field is required only if killType is set to FillOrKillDontCancel.
	//
	CancelOnInRunning bool `xml:"CancelOnInRunning,attr,omitempty"`

	//
	// The effect of this option is to cancel any unmatched bets if the selection is reset.  This can occur when the Market is reset (eg a goal is scored in an in-running market).  This defaults to true - unmatched bets will be cancelled if an event occurs in the market such that Betdaq forces the market to be reset
	// This field is required only if killType is set to FillOrKillDontCancel.
	//
	CancelIfSelectionReset bool `xml:"CancelIfSelectionReset,attr,omitempty"`

	// This field is required only if killType is set to FillOrKillDontCancel. Reprice(1), Cancel(2), DontReprice(3)
	WithdrawalRepriceOption byte `xml:"WithdrawalRepriceOption,attr,omitempty"`

	//
	// An expires at value set in the past will cause the bet to be cancelled - although the bet status will not immediately be set to Cancelled (this will occur in several moments on the exchange), the bet will not be available for matching.
	// This field is optional if killType is set to FillOrKillDontCancel, otherwise not needed at all.
	//
	ExpiresAt time.Time `xml:"ExpiresAt,attr,omitempty"`

	// Deprecated - This field is required only if killType is set to FillOrKillDontCancel.
	RestrictOrderToBroker bool `xml:"RestrictOrderToBroker,attr,omitempty"`

	// For internal use only
	ChannelTypeInfo string `xml:"ChannelTypeInfo,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`
}

type ArrayOfInt struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ArrayOfInt"`

	Int []int32 `xml:"int,omitempty"`
}

type GetAccountBalancesRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetAccountBalancesRequest"`
}

type UpdateOrdersNoReceiptRequestItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UpdateOrdersNoReceiptRequestItem"`

	BetId int64 `xml:"BetId,attr,omitempty"`

	DeltaStake float64 `xml:"DeltaStake,attr,omitempty"`

	Price float64 `xml:"Price,attr,omitempty"`

	ExpectedSelectionResetCount byte `xml:"ExpectedSelectionResetCount,attr,omitempty"`

	ExpectedWithdrawalSequenceNumber byte `xml:"ExpectedWithdrawalSequenceNumber,attr,omitempty"`

	CancelOnInRunning bool `xml:"CancelOnInRunning,attr,omitempty"`

	CancelIfSelectionReset bool `xml:"CancelIfSelectionReset,attr,omitempty"`

	SetToBeSPIfUnmatched bool `xml:"SetToBeSPIfUnmatched,attr,omitempty"`
}

type UpdateOrdersNoReceiptResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UpdateOrdersNoReceiptResponseItem"`

	BetId int64 `xml:"BetId,attr,omitempty"`

	ReturnCode int32 `xml:"ReturnCode,attr,omitempty"`
}

type UpdateOrdersNoReceiptRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UpdateOrdersNoReceiptRequest"`

	Orders struct {
		Order []*UpdateOrdersNoReceiptRequestItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type UpdateOrdersNoReceiptResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UpdateOrdersNoReceiptResponse"`

	*BaseResponse

	Orders struct {
		Order []*UpdateOrdersNoReceiptResponseItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`
}

type ListAccountPostingsRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsRequest"`

	StartTime time.Time `xml:"StartTime,attr,omitempty"`

	EndTime time.Time `xml:"EndTime,attr,omitempty"`
}

type ListAccountPostingsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsResponse"`

	*BaseResponse

	Orders struct {
		Order []*ListAccountPostingsResponseItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`

	Currency string `xml:"Currency,attr,omitempty"`

	AvailableFunds float64 `xml:"AvailableFunds,attr,omitempty"`

	Balance float64 `xml:"Balance,attr,omitempty"`

	Credit float64 `xml:"Credit,attr,omitempty"`

	Exposure float64 `xml:"Exposure,attr,omitempty"`

	HaveAllPostingsBeenReturned bool `xml:"HaveAllPostingsBeenReturned,attr,omitempty"`
}

type ListAccountPostingsResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsResponseItem"`

	PostedAt time.Time `xml:"PostedAt,attr,omitempty"`

	Description string `xml:"Description,attr,omitempty"`

	Amount float64 `xml:"Amount,attr,omitempty"`

	ResultingBalance float64 `xml:"ResultingBalance,attr,omitempty"`

	PostingCategory byte `xml:"PostingCategory,attr,omitempty"`

	OrderId int64 `xml:"OrderId,attr,omitempty"`

	MarketId int64 `xml:"MarketId,attr,omitempty"`

	TransactionId int64 `xml:"TransactionId,attr,omitempty"`
}

type ListAccountPostingsByIdRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsByIdRequest"`

	TransactionId int64 `xml:"TransactionId,attr,omitempty"`
}

type ListAccountPostingsByIdResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsByIdResponse"`

	*BaseResponse

	Orders struct {
		Order []*ListAccountPostingsByIdResponseItem `xml:"Order,omitempty"`
	} `xml:"Orders,omitempty"`

	Currency string `xml:"Currency,attr,omitempty"`

	AvailableFunds float64 `xml:"AvailableFunds,attr,omitempty"`

	Balance float64 `xml:"Balance,attr,omitempty"`

	Credit float64 `xml:"Credit,attr,omitempty"`

	Exposure float64 `xml:"Exposure,attr,omitempty"`
}

type ListAccountPostingsByIdResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListAccountPostingsByIdResponseItem"`

	PostedAt time.Time `xml:"PostedAt,attr,omitempty"`

	Description string `xml:"Description,attr,omitempty"`

	Amount float64 `xml:"Amount,attr,omitempty"`

	ResultingBalance float64 `xml:"ResultingBalance,attr,omitempty"`

	PostingCategory byte `xml:"PostingCategory,attr,omitempty"`

	OrderId int64 `xml:"OrderId,attr,omitempty"`

	MarketId int64 `xml:"MarketId,attr,omitempty"`

	TransactionId int64 `xml:"TransactionId,attr,omitempty"`
}

type GetOrderDetailsRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOrderDetailsRequest"`

	OrderId int64 `xml:"OrderId,attr,omitempty"`
}

type GetOrderDetailsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOrderDetailsResponse"`

	*BaseResponse

	OrderSettlementInformation *OrderSettlementInformationType `xml:"OrderSettlementInformation,omitempty"`

	AuditLog struct {
		AuditLog []*AuditLogItem `xml:"AuditLog,omitempty"`
	} `xml:"AuditLog,omitempty"`

	SelectionId int64 `xml:"SelectionId,attr,omitempty"`

	// Unmatched(1), Matched(2), Cancelled(3), Settled(4), Void(5), Suspended(6)
	OrderStatus byte `xml:"OrderStatus,attr,omitempty"`

	IssuedAt time.Time `xml:"IssuedAt,attr,omitempty"`

	LastChangedAt time.Time `xml:"LastChangedAt,attr,omitempty"`

	ExpiresAt time.Time `xml:"ExpiresAt,attr,omitempty"`

	ValidFrom time.Time `xml:"ValidFrom,attr,omitempty"`

	RestrictOrderToBroker bool `xml:"RestrictOrderToBroker,attr,omitempty"`

	// Normal(1), FillAndKill(2), FillOrKill(3), FillOrKillDontCancel(4)
	OrderFillType byte `xml:"OrderFillType,attr,omitempty"`

	FillOrKillThreshold float64 `xml:"FillOrKillThreshold,attr,omitempty"`

	MarketId int64 `xml:"MarketId,attr,omitempty"`

	// Inactive(1), Active(2), Suspended(3), Completed(4), Settled(6), Voided(7)
	MarketStatus byte `xml:"MarketStatus,attr,omitempty"`

	RequestedStake float64 `xml:"RequestedStake,attr,omitempty"`

	RequestedPrice float64 `xml:"RequestedPrice,attr,omitempty"`

	ExpectedSelectionResetCount byte `xml:"ExpectedSelectionResetCount,attr,omitempty"`

	TotalStake float64 `xml:"TotalStake,attr,omitempty"`

	UnmatchedStake float64 `xml:"UnmatchedStake,attr,omitempty"`

	AveragePrice float64 `xml:"AveragePrice,attr,omitempty"`

	MatchingTimeStamp time.Time `xml:"MatchingTimeStamp,attr,omitempty"`

	Polarity byte `xml:"Polarity,attr,omitempty"`

	// Reprice(1), Cancel(2), DontReprice(3)
	WithdrawalRepriceOption byte `xml:"WithdrawalRepriceOption,attr,omitempty"`

	CancelOnInRunning bool `xml:"CancelOnInRunning,attr,omitempty"`

	CancelIfSelectionReset bool `xml:"CancelIfSelectionReset,attr,omitempty"`

	SequenceNumber int64 `xml:"SequenceNumber,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`

	// Win (1), Place (2), MatchOdds (3), OverUnder (4), AsianHandicap (10), TwoBall (11), ThreeBall (12), Unspecified (13), MatchMarket (14), SetMarket (15), Moneyline (16), Total (17), Handicap (18), EachWayNonHandicap (19), EachWayHandicap (20), EachWayTournament (21), RunningBall (22),  MatchBetting (23), MatchBettingInclDraw (24), CorrectScore (25), HalfTimeFullTime (26), TotalGoals (27), GoalsScored (28), Corners (29), OddsOrEvens (30), HalfTimeResult (31), HalfTimeScore (32), MatchOddsExtraTime (33), CorrectScoreExtraTime (34), OverUnderExtraTime (35), ToQualify (36), DrawNoBet (37), HalftimeAsianHcp (39), HalftimeOverUnder (40), NextGoal (41), FirstGoalscorer (42), LastGoalscorer (43), PlayerToScore (44), FirstHalfHandicap (45), FirstHalfTotal (46), SetBetting (47), GroupBetting (48), MatchplaySingle (49), MatchplayFourball (50), MatchplayFoursome (51), TiedMatch (52), TopBatsman (53), InningsRuns (54), TotalTries (55), TotalPoints (56), FrameBetting (57), ToScoreFirst (58), ToScoreLast (59), FirstScoringPlay (60), LastScoringPlay (61), HighestScoringQtr (62), RunLine (63), RoundBetting (64), LineBetting (65)
	MarketType byte `xml:"MarketType,attr,omitempty"`

	ExpectedWithdrawalSequenceNumber byte `xml:"ExpectedWithdrawalSequenceNumber,attr,omitempty"`
}

type OrderSettlementInformationType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ OrderSettlementInformationType"`

	GrossSettlementAmount float64 `xml:"GrossSettlementAmount,attr,omitempty"`

	OrderCommission float64 `xml:"OrderCommission,attr,omitempty"`

	MarketCommission float64 `xml:"MarketCommission,attr,omitempty"`

	MarketSettledDate time.Time `xml:"MarketSettledDate,attr,omitempty"`
}

type AuditLogItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ AuditLogItem"`

	MatchedOrderInformation *MatchedOrderInformationType `xml:"MatchedOrderInformation,omitempty"`
	CommissionInformation   *CommissionInformationType   `xml:"CommissionInformation,omitempty"`

	Time time.Time `xml:"Time,attr,omitempty"`

	// Placed(1), ExplicitlyUpdated(2), Matched(3), CancelledExplicitly(4), CancelledByReset(5), CancelledOnInRunning(6), Expired(7), MatchedPortionRepricedByR4(8), UnmatchedPortionRepricedByR4(9), UnmatchedPortionCancelledByWithdrawal(10), Voided(11), Settled(12), Suspended(13), Unsuspended(14)
	OrderActionType byte `xml:"OrderActionType,attr,omitempty"`

	RequestedStake float64 `xml:"RequestedStake,attr,omitempty"`

	TotalStake float64 `xml:"TotalStake,attr,omitempty"`

	TotalAgainstStake float64 `xml:"TotalAgainstStake,attr,omitempty"`

	RequestedPrice float64 `xml:"RequestedPrice,attr,omitempty"`

	AveragePrice float64 `xml:"AveragePrice,attr,omitempty"`
}

type MatchedOrderInformationType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ MatchedOrderInformationType"`

	MatchedStake float64 `xml:"MatchedStake,attr,omitempty"`

	MatchedAgainstStake float64 `xml:"MatchedAgainstStake,attr,omitempty"`

	PriceMatched float64 `xml:"PriceMatched,attr,omitempty"`

	MatchedOrderID int64 `xml:"MatchedOrderID,attr,omitempty"`

	WasMake bool `xml:"WasMake,attr,omitempty"`
}

type CommissionInformationType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CommissionInformationType"`

	GrossSettlementAmount float64 `xml:"GrossSettlementAmount,attr,omitempty"`

	OrderCommission float64 `xml:"OrderCommission,attr,omitempty"`
}

type ListTopLevelEventsRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListTopLevelEventsRequest"`

	WantPlayMarkets bool `xml:"WantPlayMarkets,attr,omitempty"`
}

type ListTopLevelEventsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListTopLevelEventsResponse"`

	*BaseResponse

	EventClassifiers *EventClassifierType `xml:"EventClassifiers,omitempty"`
}

type EventClassifierType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ EventClassifierType"`

	EventClassifiers []*EventClassifierType `xml:"EventClassifiers,omitempty"`
	Markets          []*MarketType          `xml:"Markets,omitempty"`

	Id int64 `xml:"Id,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	DisplayOrder int16 `xml:"DisplayOrder,attr,omitempty"`

	IsEnabledForMultiples bool `xml:"IsEnabledForMultiples,attr,omitempty"`

	ParentId int64 `xml:"ParentId,attr,omitempty"`
}

type GetEventSubTreeWithSelectionsRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeWithSelectionsRequest"`

	EventClassifierIds int64 `xml:"EventClassifierIds,omitempty"`

	WantPlayMarkets bool `xml:"WantPlayMarkets,attr,omitempty"`
}

type GetEventSubTreeWithSelectionsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeWithSelectionsResponse"`

	*BaseResponse

	EventClassifiers *EventClassifierType `xml:"EventClassifiers,omitempty"`
}

type GetEventSubTreeNoSelectionsRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeNoSelectionsRequest"`

	EventClassifierIds int64 `xml:"EventClassifierIds,omitempty"`

	WantDirectDescendentsOnly bool `xml:"WantDirectDescendentsOnly,attr,omitempty"`

	WantPlayMarkets bool `xml:"WantPlayMarkets,attr,omitempty"`
}

type GetEventSubTreeNoSelectionsResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetEventSubTreeNoSelectionsResponse"`

	*BaseResponse

	EventClassifiers *EventClassifierType `xml:"EventClassifiers,omitempty"`
}

type MarketType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ MarketType"`

	Selections []*SelectionType `xml:"Selections,omitempty"`

	Id int64 `xml:"Id,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	Type int16 `xml:"Type,attr,omitempty"`

	IsPlayMarket bool `xml:"IsPlayMarket,attr,omitempty"`

	// Inactive(1), Active(2), Suspended(3), Withdrawn(4), BallotedOut(9), Voided(5), Completed(6), Settled(8)
	Status int16 `xml:"Status,attr,omitempty"`

	NumberOfWinningSelections int16 `xml:"NumberOfWinningSelections,attr,omitempty"`

	StartTime time.Time `xml:"StartTime,attr,omitempty"`

	WithdrawalSequenceNumber int16 `xml:"WithdrawalSequenceNumber,attr,omitempty"`

	DisplayOrder int16 `xml:"DisplayOrder,attr,omitempty"`

	IsEnabledForMultiples bool `xml:"IsEnabledForMultiples,attr,omitempty"`

	IsInRunningAllowed bool `xml:"IsInRunningAllowed,attr,omitempty"`

	RaceGrade string `xml:"RaceGrade,attr,omitempty"`

	IsManagedWhenInRunning bool `xml:"IsManagedWhenInRunning,attr,omitempty"`

	IsCurrentlyInRunning bool `xml:"IsCurrentlyInRunning,attr,omitempty"`

	InRunningDelaySeconds int32 `xml:"InRunningDelaySeconds,attr,omitempty"`

	EventClassifierId int64 `xml:"EventClassifierId,attr,omitempty"`

	PlacePayout float64 `xml:"PlacePayout,attr,omitempty"`
}

type SelectionType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SelectionType"`

	Id int64 `xml:"Id,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	Status int16 `xml:"Status,attr,omitempty"`

	ResetCount int16 `xml:"ResetCount,attr,omitempty"`

	DeductionFactor float64 `xml:"DeductionFactor,attr,omitempty"`

	DisplayOrder int32 `xml:"DisplayOrder,attr,omitempty"`
}

type GetMarketInformationRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetMarketInformationRequest"`
}

type GetMarketInformationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetMarketInformationResponse"`

	*BaseResponse

	Markets *MarketType `xml:"Markets,omitempty"`
}

type ChangePasswordResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangePasswordResponse"`

	*BaseResponse
}

type ChangePasswordRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangePasswordRequest"`

	Password string `xml:"Password,attr,omitempty"`
}

type ListMarketWithdrawalHistoryResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListMarketWithdrawalHistoryResponse"`

	*BaseResponse

	Withdrawals *ListMarketWithdrawalHistoryResponseItem `xml:"Withdrawals,omitempty"`
}

type ListMarketWithdrawalHistoryResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListMarketWithdrawalHistoryResponseItem"`

	SelectionId int64 `xml:"SelectionId,attr,omitempty"`

	WithdrawalTime time.Time `xml:"WithdrawalTime,attr,omitempty"`

	SequenceNumber int16 `xml:"SequenceNumber,attr,omitempty"`

	ReductionFactor float64 `xml:"ReductionFactor,attr,omitempty"`

	CompoundReductionFactor float64 `xml:"CompoundReductionFactor,attr,omitempty"`
}

type ListMarketWithdrawalHistoryRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListMarketWithdrawalHistoryRequest"`

	MarketId int64 `xml:"MarketId,attr,omitempty"`
}

type GetPricesRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetPricesRequest"`

	MarketIds int64 `xml:"MarketIds,omitempty"`

	ThresholdAmount float64 `xml:"ThresholdAmount,attr,omitempty"`

	NumberForPricesRequired int32 `xml:"NumberForPricesRequired,attr,omitempty"`

	NumberAgainstPricesRequired int32 `xml:"NumberAgainstPricesRequired,attr,omitempty"`

	WantMarketMatchedAmount bool `xml:"WantMarketMatchedAmount,attr,omitempty"`

	WantSelectionsMatchedAmounts bool `xml:"WantSelectionsMatchedAmounts,attr,omitempty"`

	WantSelectionMatchedDetails bool `xml:"WantSelectionMatchedDetails,attr,omitempty"`
}

type GetPricesResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetPricesResponse"`

	*BaseResponse

	MarketPrices *MarketTypeWithPrices `xml:"MarketPrices,omitempty"`
}

type SelectionTypeWithPrices struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SelectionTypeWithPrices"`

	ForSidePrices     *PricesType `xml:"ForSidePrices,omitempty"`
	AgainstSidePrices *PricesType `xml:"AgainstSidePrices,omitempty"`

	Id int64 `xml:"Id,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	Status int16 `xml:"Status,attr,omitempty"`

	ResetCount int16 `xml:"ResetCount,attr,omitempty"`

	DeductionFactor float64 `xml:"DeductionFactor,attr,omitempty"`

	MatchedSelectionForStake float64 `xml:"MatchedSelectionForStake,attr,omitempty"`

	SelectionOpenInterest float64 `xml:"SelectionOpenInterest,attr,omitempty"`

	MarketWinnings float64 `xml:"MarketWinnings,attr,omitempty"`

	MarketPositiveWinnings float64 `xml:"MarketPositiveWinnings,attr,omitempty"`

	MatchedSelectionAgainstStake float64 `xml:"MatchedSelectionAgainstStake,attr,omitempty"`

	LastMatchedOccurredAt time.Time `xml:"LastMatchedOccurredAt,attr,omitempty"`

	LastMatchedPrice float64 `xml:"LastMatchedPrice,attr,omitempty"`

	LastMatchedForSideAmount float64 `xml:"LastMatchedForSideAmount,attr,omitempty"`

	LastMatchedAgainstSideAmount float64 `xml:"LastMatchedAgainstSideAmount,attr,omitempty"`

	MatchedForSideAmountAtSamePrice float64 `xml:"MatchedForSideAmountAtSamePrice,attr,omitempty"`

	MatchedAgainstSideAmountAtSamePrice float64 `xml:"MatchedAgainstSideAmountAtSamePrice,attr,omitempty"`

	FirstMatchAtSamePriceOccurredAt time.Time `xml:"FirstMatchAtSamePriceOccurredAt,attr,omitempty"`

	NumberOrders int32 `xml:"NumberOrders,attr,omitempty"`

	NumberPunters int32 `xml:"NumberPunters,attr,omitempty"`
}

type MarketTypeWithPrices struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ MarketTypeWithPrices"`

	Selections []*SelectionTypeWithPrices `xml:"Selections,omitempty"`

	Id int64 `xml:"Id,attr,omitempty"`

	Name string `xml:"Name,attr,omitempty"`

	Type int16 `xml:"Type,attr,omitempty"`

	IsPlayMarket bool `xml:"IsPlayMarket,attr,omitempty"`

	Status int16 `xml:"Status,attr,omitempty"`

	NumberOfWinningSelections int16 `xml:"NumberOfWinningSelections,attr,omitempty"`

	StartTime time.Time `xml:"StartTime,attr,omitempty"`

	WithdrawalSequenceNumber int16 `xml:"WithdrawalSequenceNumber,attr,omitempty"`

	DisplayOrder int16 `xml:"DisplayOrder,attr,omitempty"`

	IsEnabledForMultiples bool `xml:"IsEnabledForMultiples,attr,omitempty"`

	IsInRunningAllowed bool `xml:"IsInRunningAllowed,attr,omitempty"`

	IsManagedWhenInRunning bool `xml:"IsManagedWhenInRunning,attr,omitempty"`

	IsCurrentlyInRunning bool `xml:"IsCurrentlyInRunning,attr,omitempty"`

	InRunningDelaySeconds int32 `xml:"InRunningDelaySeconds,attr,omitempty"`

	ReturnCode int32 `xml:"ReturnCode,attr,omitempty"`

	TotalMatchedAmount float64 `xml:"TotalMatchedAmount,attr,omitempty"`

	PlacePayout float64 `xml:"PlacePayout,attr,omitempty"`

	MatchedMarketForStake float64 `xml:"MatchedMarketForStake,attr,omitempty"`

	MatchedMarketAgainstStake float64 `xml:"MatchedMarketAgainstStake,attr,omitempty"`

	HomeTeamScore int32 `xml:"HomeTeamScore,attr,omitempty"`

	AwayTeamScore int32 `xml:"AwayTeamScore,attr,omitempty"`

	ScoreType int16 `xml:"ScoreType,attr,omitempty"`
}

type PricesType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PricesType"`

	Price float64 `xml:"Price,attr,omitempty"`

	Stake float64 `xml:"Stake,attr,omitempty"`
}

type CancelAllOrdersOnMarketRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersOnMarketRequest"`

	MarketIds int64 `xml:"MarketIds,omitempty"`
}

type CancelAllOrdersOnMarketResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersOnMarketResponseItem"`

	OrderHandle int64 `xml:"OrderHandle,attr,omitempty"`

	CancelledForSideStake float64 `xml:"cancelledForSideStake,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`
}

type CancelAllOrdersOnMarketResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersOnMarketResponse"`

	*BaseResponse

	OrderIds int64                                `xml:"OrderIds,omitempty"`
	Order    *CancelAllOrdersOnMarketResponseItem `xml:"Order,omitempty"`
}

type CancelAllOrdersResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ CancelAllOrdersResponseItem"`

	OrderHandle int64 `xml:"OrderHandle,attr,omitempty"`

	CancelledForSideStake float64 `xml:"cancelledForSideStake,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`
}

type SuspendFromTradingRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendFromTradingRequest"`
}

type SuspendFromTradingResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendFromTradingResponse"`

	*BaseResponse
}

type UnsuspendFromTradingRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendFromTradingRequest"`
}

type UnsuspendFromTradingResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendFromTradingResponse"`

	*BaseResponse
}

type SuspendOrdersResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendOrdersResponseItem"`

	OrderId int64 `xml:"OrderId,attr,omitempty"`

	SuspendedForSideStake float64 `xml:"SuspendedForSideStake,attr,omitempty"`

	PunterReferenceNumber int64 `xml:"PunterReferenceNumber,attr,omitempty"`
}

type SuspendOrdersRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendOrdersRequest"`

	OrderIds int64 `xml:"OrderIds,omitempty"`
}

type SuspendOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendOrdersResponse"`

	*BaseResponse

	Orders *SuspendOrdersResponseItem `xml:"Orders,omitempty"`
}

type SuspendAllOrdersOnMarketRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrdersOnMarketRequest"`

	MarketId int64 `xml:"MarketId,attr,omitempty"`
}

type SuspendAllOrdersOnMarketResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrdersOnMarketResponse"`

	*BaseResponse

	Orders *SuspendOrdersResponseItem `xml:"Orders,omitempty"`
}

type UnsuspendOrdersRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendOrdersRequest"`

	OrderIds int64 `xml:"OrderIds,omitempty"`
}

type UnsuspendOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ UnsuspendOrdersResponse"`

	*BaseResponse
}

type SuspendAllOrdersRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrdersRequest"`
}

type SuspendAllOrdersResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SuspendAllOrdersResponse"`

	*BaseResponse

	Orders *SuspendOrdersResponseItem `xml:"Orders,omitempty"`
}

type ListBlacklistInformationRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBlacklistInformationRequest"`
}

type ListBlacklistInformationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListBlacklistInformationResponse"`

	*BaseResponse

	ApiNamesAndTimes *ApiTimes `xml:"ApiNamesAndTimes,omitempty"`
}

type ApiTimes struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ApiTimes"`

	ApiName string `xml:"ApiName,attr,omitempty"`

	RemainingMS int32 `xml:"RemainingMS,attr,omitempty"`
}

type RegisterHeartbeatRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ RegisterHeartbeatRequest"`

	ThresholdMs int32 `xml:"ThresholdMs,attr,omitempty"`

	// CancelOrders(1), SuspendOrders(2), SuspendPunter(3)
	HeartbeatAction byte `xml:"HeartbeatAction,attr,omitempty"`
}

type RegisterHeartbeatResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ RegisterHeartbeatResponse"`

	*BaseResponse
}

type ChangeHeartbeatRegistrationRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangeHeartbeatRegistrationRequest"`

	ThresholdMs int32 `xml:"ThresholdMs,attr,omitempty"`

	// CancelOrders(1), SuspendOrders(2), SuspendPunter(3)
	HeartbeatAction byte `xml:"HeartbeatAction,attr,omitempty"`
}

type ChangeHeartbeatRegistrationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ChangeHeartbeatRegistrationResponse"`

	*BaseResponse
}

type DeregisterHeartbeatRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ DeregisterHeartbeatRequest"`
}

type DeregisterHeartbeatResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ DeregisterHeartbeatResponse"`

	*BaseResponse
}

type PulseRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PulseRequest"`
}

type PulseResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ PulseResponse"`

	*BaseResponse

	PerformedAt time.Time `xml:"PerformedAt,attr,omitempty"`

	// CancelOrders(1), SuspendOrders(2), SuspendPunter(3)
	HeartbeatAction byte `xml:"HeartbeatAction,attr,omitempty"`
}

type ExternalApiHeader struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ExternalApiHeader"`

	Version float64 `xml:"version,attr,omitempty"`

	LanguageCode string `xml:"languageCode,attr,omitempty"`

	Username string `xml:"username,attr,omitempty"`

	Password string `xml:"password,attr,omitempty"`

	ApplicationIdentifier string `xml:"applicationIdentifier,attr,omitempty"`
}

type GetOddsLadderRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadderRequest"`

	PriceFormat byte `xml:"PriceFormat,attr,omitempty"`
}

type GetOddsLadderResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadderResponse"`

	*BaseResponse

	Ladder []*GetOddsLadderResponseItem `xml:"Ladder,omitempty"`
}

type GetOddsLadderResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetOddsLadderResponseItem"`

	Price float64 `xml:"price,attr,omitempty"`

	Representation string `xml:"representation,attr,omitempty"`
}

type GetCurrentSelectionSequenceNumberRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetCurrentSelectionSequenceNumberRequest"`
}

type GetCurrentSelectionSequenceNumberResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetCurrentSelectionSequenceNumberResponse"`

	*BaseResponse

	SelectionSequenceNumber int64 `xml:"SelectionSequenceNumber,attr,omitempty"`
}

type ListSelectionTradesRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionTradesRequest"`

	SelectionRequests *SelectionTradesRequestItem `xml:"selectionRequests,omitempty"`

	Currency string `xml:"currency,attr,omitempty"`
}

type SelectionTradesRequestItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ SelectionTradesRequestItem"`

	SelectionId int64 `xml:"selectionId,attr,omitempty"`

	FromTradeId int64 `xml:"fromTradeId,attr,omitempty"`
}

type ListSelectionTradesResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionTradesResponse"`

	*BaseResponse

	SelectionTrades *ListSelectionTradesResponseItem `xml:"SelectionTrades,omitempty"`
}

type ListSelectionTradesResponseItem struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ ListSelectionTradesResponseItem"`

	TradeItems *TradeItemType `xml:"TradeItems,omitempty"`

	SelectionId int64 `xml:"selectionId,attr,omitempty"`

	MaxTradeId int64 `xml:"maxTradeId,attr,omitempty"`

	MaxTradeIdReturned int64 `xml:"maxTradeIdReturned,attr,omitempty"`
}

type TradeItemType struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ TradeItemType"`

	OccurredAt time.Time `xml:"occurredAt,attr,omitempty"`

	Price float64 `xml:"price,attr,omitempty"`

	BackersStake float64 `xml:"backersStake,attr,omitempty"`

	LayersLiability float64 `xml:"layersLiability,attr,omitempty"`

	// Back(1), Lay(2)
	TradeType byte `xml:"tradeType,attr,omitempty"`
}

type GetSPEnabledMarketsInformationRequest struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetSPEnabledMarketsInformationRequest"`
}

type GetSPEnabledMarketsInformationResponse struct {
	XMLName xml.Name `xml:"http://www.GlobalBettingExchange.com/ExternalAPI/ GetSPEnabledMarketsInformationResponse"`

	*BaseResponse

	SPEnabledEvent struct {
		MarketTypeIds struct {
			MarketTypeId []byte `xml:"MarketTypeId,omitempty"`
		} `xml:"MarketTypeIds,omitempty"`

		EventId int64 `xml:"eventId,attr,omitempty"`
	} `xml:"SPEnabledEvent,omitempty"`
}

type ReadOnlyService struct {
	client *SOAPClient
}

func NewReadOnlyService(url string, tls bool, auth *BasicAuth) *ReadOnlyService {
	if url == "" {
		url = "http://api.betdaq.com/v2.0/ReadOnlyService.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &ReadOnlyService{
		client: client,
	}
}

/* Returns the selections with a sequence number greater than the sequence number provided. */
func (service *ReadOnlyService) ListSelectionsChangedSince(request *ListSelectionsChangedSince) (*ListSelectionsChangedSinceResponse, error) {
	response := new(ListSelectionsChangedSinceResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListSelectionsChangedSince", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the set of top level events that are currently active. */
func (service *ReadOnlyService) ListTopLevelEvents(request *ListTopLevelEvents) (*ListTopLevelEventsResponse, error) {
	response := new(ListTopLevelEventsResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListTopLevelEvents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the tree of events and markets. */
func (service *ReadOnlyService) GetEventSubTreeWithSelections(request *GetEventSubTreeWithSelections) (*GetEventSubTreeWithSelectionsResponse, error) {
	response := new(GetEventSubTreeWithSelectionsResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetEventSubTreeWithSelections", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the tree of events and markets. */
func (service *ReadOnlyService) GetEventSubTreeNoSelections(request *GetEventSubTreeNoSelections) (*GetEventSubTreeNoSelectionsResponse, error) {
	response := new(GetEventSubTreeNoSelectionsResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetEventSubTreeNoSelections", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the tree of events and markets. */
func (service *ReadOnlyService) GetMarketInformation(request *GetMarketInformation) (*GetMarketInformationResponse, error) {
	response := new(GetMarketInformationResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetMarketInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns history of withdrawals for a given market. */
func (service *ReadOnlyService) ListMarketWithdrawalHistory(request *ListMarketWithdrawalHistory) (*ListMarketWithdrawalHistoryResponse, error) {
	response := new(ListMarketWithdrawalHistoryResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListMarketWithdrawalHistory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the prices for the specified markets. */
func (service *ReadOnlyService) GetPrices(request *GetPrices) (*GetPricesResponse, error) {
	response := new(GetPricesResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetPrices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the current odds ladder in requested price format." */
func (service *ReadOnlyService) GetOddsLadder(request *GetOddsLadder) (*GetOddsLadderResponse, error) {
	response := new(GetOddsLadderResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetOddsLadder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the maximum selection sequence number. */
func (service *ReadOnlyService) GetCurrentSelectionSequenceNumber(request *GetCurrentSelectionSequenceNumber) (*GetCurrentSelectionSequenceNumberResponse, error) {
	response := new(GetCurrentSelectionSequenceNumberResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetCurrentSelectionSequenceNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns the history of trades on the selection(s) specified. */
func (service *ReadOnlyService) ListSelectionTrades(request *ListSelectionTrades) (*ListSelectionTradesResponse, error) {
	response := new(ListSelectionTradesResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListSelectionTrades", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns information defining which markets are enabled for starting-price orders. */
func (service *ReadOnlyService) GetSPEnabledMarketsInformation(request *GetSPEnabledMarketsInformation) (*GetSPEnabledMarketsInformationResponse, error) {
	response := new(GetSPEnabledMarketsInformationResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetSPEnabledMarketsInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type SecureService struct {
	client *SOAPClient
}

func NewSecureService(url string, tls bool, auth *BasicAuth) *SecureService {
	if url == "" {
		url = "https://api.betdaq.com/v2.0/Secure/SecureService.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &SecureService{
		client: client,
	}
}

/* Returns bootstrap orders that have a sequence number greater than the sequence number specified. */
func (service *SecureService) ListBootstrapOrders(request *ListBootstrapOrders) (*ListBootstrapOrdersResponse, error) {
	response := new(ListBootstrapOrdersResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListBootstrapOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns a list of orders that have changed since a given sequence number. */
func (service *SecureService) ListOrdersChangedSince(request *ListOrdersChangedSince) (*ListOrdersChangedSinceResponse, error) {
	response := new(ListOrdersChangedSinceResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListOrdersChangedSince", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Places one or more orders on the exchange. */
func (service *SecureService) PlaceOrdersNoReceipt(request *PlaceOrdersNoReceipt) (*PlaceOrdersNoReceiptResponse, error) {
	response := new(PlaceOrdersNoReceiptResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/PlaceOrdersNoReceipt", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Places one or more orders on the exchange and waits for response from exchange. */
func (service *SecureService) PlaceOrdersWithReceipt(request *PlaceOrdersWithReceipt) (*PlaceOrdersWithReceiptResponse, error) {
	response := new(PlaceOrdersWithReceiptResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/PlaceOrdersWithReceipt", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Cancels one or more orders on the exchange. */
func (service *SecureService) CancelOrders(request *CancelOrders) (*CancelOrdersResponse, error) {
	response := new(CancelOrdersResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/CancelOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Cancels all unmatched orders across all markets. */
func (service *SecureService) CancelAllOrders(request *CancelAllOrders) (*CancelAllOrdersResponse, error) {
	response := new(CancelAllOrdersResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/CancelAllOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Cancels all unmatched orders for specified markets. */
func (service *SecureService) CancelAllOrdersOnMarket(request *CancelAllOrdersOnMarket) (*CancelAllOrdersOnMarketResponse, error) {
	response := new(CancelAllOrdersOnMarketResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/CancelAllOrdersOnMarket", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns an account summary for the current punter. */
func (service *SecureService) GetAccountBalances(request *GetAccountBalances) (*GetAccountBalancesResponse, error) {
	response := new(GetAccountBalancesResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetAccountBalances", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns an account summary and details of orders placed for the current punter. */
func (service *SecureService) ListAccountPostings(request *ListAccountPostings) (*ListAccountPostingsResponse, error) {
	response := new(ListAccountPostingsResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListAccountPostings", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns an account summary and details of orders placed for the current punter. */
func (service *SecureService) ListAccountPostingsById(request *ListAccountPostingsById) (*ListAccountPostingsByIdResponse, error) {
	response := new(ListAccountPostingsByIdResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListAccountPostingsById", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Updates one or more orders on the exchange without waiting for a response. */
func (service *SecureService) UpdateOrdersNoReceipt(request *UpdateOrdersNoReceipt) (*UpdateOrdersNoReceiptResponse, error) {
	response := new(UpdateOrdersNoReceiptResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/UpdateOrdersNoReceipt", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Gets detailed information about an order. */
func (service *SecureService) GetOrderDetails(request *GetOrderDetails) (*GetOrderDetailsResponse, error) {
	response := new(GetOrderDetailsResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/GetOrderDetails", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Change password on the current account. */
func (service *SecureService) ChangePassword(request *ChangePassword) (*ChangePasswordResponse, error) {
	response := new(ChangePasswordResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ChangePassword", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Suspend any of your orders from being matched. */
func (service *SecureService) SuspendFromTrading(request *SuspendFromTrading) (*SuspendFromTradingResponse, error) {
	response := new(SuspendFromTradingResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/SuspendFromTrading", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Unsuspend yourself from being suspending from trading. */
func (service *SecureService) UnsuspendFromTrading(request *UnsuspendFromTrading) (*UnsuspendFromTradingResponse, error) {
	response := new(UnsuspendFromTradingResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/UnsuspendFromTrading", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Suspends one or more Orders. */
func (service *SecureService) SuspendOrders(request *SuspendOrders) (*SuspendOrdersResponse, error) {
	response := new(SuspendOrdersResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/SuspendOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Suspends all unmatched orders on a market. */
func (service *SecureService) SuspendAllOrdersOnMarket(request *SuspendAllOrdersOnMarket) (*SuspendAllOrdersOnMarketResponse, error) {
	response := new(SuspendAllOrdersOnMarketResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/SuspendAllOrdersOnMarket", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Unsuspends one or more suspended Orders. */
func (service *SecureService) UnsuspendOrders(request *UnsuspendOrders) (*UnsuspendOrdersResponse, error) {
	response := new(UnsuspendOrdersResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/UnsuspendOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Suspends one or more Orders. */
func (service *SecureService) SuspendAllOrders(request *SuspendAllOrders) (*SuspendAllOrdersResponse, error) {
	response := new(SuspendAllOrdersResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/SuspendAllOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* List blacklist information of the current punter. */
func (service *SecureService) ListBlacklistInformation(request *ListBlacklistInformation) (*ListBlacklistInformationResponse, error) {
	response := new(ListBlacklistInformationResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ListBlacklistInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Register HeartBeat. */
func (service *SecureService) RegisterHeartbeat(request *RegisterHeartbeat) (*RegisterHeartbeatResponse, error) {
	response := new(RegisterHeartbeatResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/RegisterHeartbeat", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Change HeartBeat registration. */
func (service *SecureService) ChangeHeartbeatRegistration(request *ChangeHeartbeatRegistration) (*ChangeHeartbeatRegistrationResponse, error) {
	response := new(ChangeHeartbeatRegistrationResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/ChangeHeartbeatRegistration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deteles HeartBeat registration. */
func (service *SecureService) DeregisterHeartbeat(request *DeregisterHeartbeat) (*DeregisterHeartbeatResponse, error) {
	response := new(DeregisterHeartbeatResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/DeregisterHeartbeat", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Pulse. */
func (service *SecureService) Pulse(request *Pulse) (*PulseResponse, error) {
	response := new(PulseResponse)
	err := service.client.Call("http://www.GlobalBettingExchange.com/ExternalAPI/Pulse", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
