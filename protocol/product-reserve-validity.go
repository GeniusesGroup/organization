/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductReserveValidity indicate the domain record data fields.
type ProductReserveValidity interface {
	ProductID() [16]byte         // product domain
	Duration() protocol.Duration // from add to invoice time
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductReserveValidity_StorageServices interface {
	Save(prv ProductReserveValidity) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (prv ProductReserveValidity, err protocol.Error)
	Last(productID [16]byte) (prv ProductReserveValidity, numbers uint64, err protocol.Error)
}
