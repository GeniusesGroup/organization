/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountTimeValidity or LiveUntil
type DiscountTimeValidity interface {
	DiscountID() [16]byte // discount domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type DiscountTimeValidity_StorageServices interface {
	Save(dtv DiscountTimeValidity) (err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dtv DiscountTimeValidity, err protocol.Error)
	Last(discountID [16]byte) (dtv DiscountTimeValidity, numbers uint64, err protocol.Error)
}
