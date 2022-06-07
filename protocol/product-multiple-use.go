/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductMultipleUse indicate the domain record data fields.
type ProductMultipleUse interface {
	ProductID() [16]byte        // product domain
	Maximum() uint64            // Maximum time this service can be served to buyer
	Timeout() protocol.Duration // from first use
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type ProductMultipleUse_StorageServices interface {
	Save(pm ProductMultipleUse) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pm ProductMultipleUse, err protocol.Error)
	Last(productID [16]byte) (pm ProductMultipleUse, numbers uint64, err protocol.Error)
}
