/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductPrice indicate the domain record data fields.
// ProductPrice can be register just by producer organization whom own the related quiddity
type ProductPrice interface {
	ProductID() [16]byte           // product domain
	Currency() [16]byte            // financial-currency
	Price() protocol.AmountOfMoney // Basic Price not payable price in the currency
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

type ProductPrice_StorageServices interface {
	Save(pp ProductPrice) (err protocol.Error)

	Count(productID [16]byte, currency uint64) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, currency uint64, versionOffset uint64) (pp ProductPrice, err protocol.Error)
	Last(productID [16]byte, currency uint64) (pp ProductPrice, numbers uint64, err protocol.Error)

	ListProductCurrencies(productID [16]byte, offset, limit uint64) (currencies []uint64, numbers uint64, err protocol.Error)
}
