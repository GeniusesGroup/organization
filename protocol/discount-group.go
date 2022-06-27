/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountGroup indicate the domain record data fields.
type DiscountGroup interface {
	DiscountID() [16]byte // discount domain
	GroupID() [16]byte    // group domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type DiscountGroup_StorageServices interface {
	Save(dg DiscountGroup) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dg DiscountGroup, numbers uint64, err protocol.Error)

	FindByGroup(groupID [16]byte, offset, limit uint64) (discountIDs [][16]byte, numbers uint64, err protocol.Error)
}
