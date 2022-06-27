/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountTimeValidity or LiveUntil indicate the domain record data fields.
type DiscountTimeValidity interface {
	DiscountID() [16]byte // discount domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type DiscountTimeValidity_StorageServices interface {
	Save(dtv DiscountTimeValidity) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dtv DiscountTimeValidity, numbers uint64, err protocol.Error)
}
