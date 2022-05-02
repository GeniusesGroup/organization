/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// OrderTrade indicate the invoice order-trade record data fields.
// can be zero to prevent sell on profit(cover order) or stopping loss
type OrderTrade interface {
	OrderID() [16]byte                 // order domain
	Level() uint8                      // usually 3 level trade order
	HighPrice() protocol.AmountOfMoney // SaveProfit	e.g. 45$ , -102$
	LowPrice() protocol.AmountOfMoney  // StopLoss		e.g. 43$ , -95$
	Type() OrderTrade_Type             //
	Time() protocol.Time               // Save time
	RequestID() [16]byte               // user-request domain
}

type OrderTrade_StorageServices interface {
	Save(ot OrderTrade) (err protocol.Error)

	Count(orderID [16]byte) (numbers uint64, err protocol.Error)
	Get(orderID [16]byte, versionOffset uint64) (ot OrderTrade, err protocol.Error)
	Last(orderID [16]byte) (ot OrderTrade, numbers uint64, err protocol.Error)
}

type OrderTrade_Type uint8

const (
	OrderTrade_Type_Unset OrderTrade_Type = iota
	// stop-loss orders guarantee execution
	OrderTrade_Type_StopLoss
	// stop-limit orders guarantee price
	OrderTrade_Type_StopLimit
)
