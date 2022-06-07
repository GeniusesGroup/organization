/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductTimeValidity indicate the domain record data fields.
// flight departure and arrival time, concert start and end time, ...
type ProductTimeValidity interface {
	ProductID() [16]byte  // product domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type ProductTimeValidity_StorageServices interface {
	Save(pt ProductTimeValidity) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTimeValidity, err protocol.Error)
	Last(productID [16]byte) (pt ProductTimeValidity, numbers uint64, err protocol.Error)
}
