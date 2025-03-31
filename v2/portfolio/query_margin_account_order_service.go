package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

type QueryMarginAccountOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	recvWindow        *int64
}

// Symbol set symbol
func (s *QueryMarginAccountOrderService) Symbol(symbol string) *QueryMarginAccountOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *QueryMarginAccountOrderService) OrderID(orderID int64) *QueryMarginAccountOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderId
func (s *QueryMarginAccountOrderService) OrigClientOrderID(origClientOrderID string) *QueryMarginAccountOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *QueryMarginAccountOrderService) RecvWindow(recvWindow int64) *QueryMarginAccountOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *QueryMarginAccountOrderService) Do(ctx context.Context) (*QueryMarginAccountOrder, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(QueryMarginAccountOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type QueryMarginAccountOrder struct {
	ClientOrderId           string      `json:"clientOrderId"`
	CummulativeQuoteQty     string      `json:"cummulativeQuoteQty"`
	ExecutedQty             string      `json:"executedQty"`
	IcebergQty              string      `json:"icebergQty"`
	IsWorking               bool        `json:"isWorking"`
	OrderId                 int         `json:"orderId"`
	OrigQty                 string      `json:"origQty"`
	Price                   string      `json:"price"`
	Side                    string      `json:"side"`
	Status                  string      `json:"status"`
	StopPrice               string      `json:"stopPrice"`
	Symbol                  string      `json:"symbol"`
	Time                    int64       `json:"time"`
	TimeInForce             string      `json:"timeInForce"`
	Type                    string      `json:"type"`
	UpdateTime              int64       `json:"updateTime"`
	AccountId               int         `json:"accountId"`
	SelfTradePreventionMode string      `json:"selfTradePreventionMode"`
	PreventedMatchId        interface{} `json:"preventedMatchId"`
	PreventedQuantity       interface{} `json:"preventedQuantity"`
}
