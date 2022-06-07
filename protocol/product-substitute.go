/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductSubstitute indicate the domain record data fields.
type ProductSubstitute interface {
	ProductID() [16]byte    // product domain
	Priority() uint64       // Substitute priority
	SubstituteID() [16]byte // product domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductSubstitute_StorageServices interface {
	Save(ps ProductSubstitute) protocol.Error

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (ps ProductSubstitute, err protocol.Error)
	Last(productID [16]byte) (ps ProductSubstitute, numbers uint64, err protocol.Error)

	FindBySubstitute(substituteID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}
