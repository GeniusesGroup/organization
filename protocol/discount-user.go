/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountUser indicate the domain record data fields.
type DiscountUser interface {
	DiscountID() [16]byte // discount domain
	UserID() [16]byte     // user domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type DiscountUser_StorageServices interface {
	Save(ds DiscountUser) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (ds DiscountUser, numbers uint64, err protocol.Error)

	FindByUser(userID [16]byte, offset, limit uint64) (discountIDs [][16]byte, numbers uint64, err protocol.Error)
}
