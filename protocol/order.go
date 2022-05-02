/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Order use to register an order to e.g. exchange(an instruction to buy or sell) currencies
// A matching engine uses the orders book to determine which orders can be fully or partially executed.
type Order interface {
	ID() [16]byte        // OrderID
	UserID() [16]byte    // user-status domain
	Amount() int64       // + || -		or Quantity
	Kind() Order_Kind    //
	Type() Order_Type    //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type Order_StorageServices interface {
	Save(o Order) (err protocol.Error)

	Count(id [16]byte) (numbers uint64, err protocol.Error)
	Get(id [16]byte, versionOffset uint64) (o Order, err protocol.Error)
	Last(id [16]byte) (o Order, numbers uint64, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}

// Use to send order to desire matching engine
type Order_Kind uint8

const (
	Order_Kind_Unset Order_Kind = 0
	Order_Kind_Exchange
	Order_Kind_Stock
)

type Order_Type uint8

const (
	Order_Type_Unset   Order_Type = 0
	Order_Type_Regular Order_Type = (1 << iota) // order-price rules
	Order_Type_Trade                            // order-trade rules
	// Whole order must be filled for matching deal or order must cancelled on timeout
	// If this flag absent, allow matching engine to deal part of order.
	// https://en.wikipedia.org/wiki/Immediate_or_cancel
	// https://en.wikipedia.org/wiki/All_or_none
	// https://en.wikipedia.org/wiki/Fill_or_kill
	Order_Type_MustFill
	Order_Type_NeedAprove
)
