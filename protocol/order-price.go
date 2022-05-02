/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// OrderPrice indicate the invoice order-price record data fields.
// An order submitted without a specified price, and sequentially matched at the best available price in the market or in other words
// if user want to limit price, must fill order-price record otherwise market price will choose for the deals.
type OrderPrice interface {
	OrderID() [16]byte             // order domain
	Price() protocol.AmountOfMoney // decimal, e.g. 1.258654, 0.987456487965
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}

type OrderPrice_StorageServices interface {
	Save(op OrderPrice) (err protocol.Error)

	Count(orderID [16]byte) (numbers uint64, err protocol.Error)
	Get(orderID [16]byte, versionOffset uint64) (op OrderPrice, err protocol.Error)
	Last(orderID [16]byte) (op OrderPrice, numbers uint64, err protocol.Error)
}
