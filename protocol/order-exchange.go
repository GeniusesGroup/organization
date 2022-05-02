/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// OrderExchange use to exchange(an instruction to buy or sell) currencies
type OrderExchange interface {
	OrderID() [16]byte       // order domain
	SellAccountID() [16]byte // financial-account || financial-bank-account domain
	BuyAccountID() [16]byte  // financial-account || financial-bank-account domain
	Time() protocol.Time     // Save time
	RequestID() [16]byte     // user-request domain
}

type OrderExchange_StorageServices interface {
	Save(oe OrderExchange) (err protocol.Error)

	Count(orderID [16]byte) (numbers uint64, err protocol.Error)
	Get(orderID [16]byte, versionOffset uint64) (oe OrderExchange, err protocol.Error)
	Last(orderID [16]byte) (oe OrderExchange, numbers uint64, err protocol.Error)
}
