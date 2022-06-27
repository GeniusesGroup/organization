/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// DiscountUserType indicate the domain record data fields.
type DiscountUserType interface {
	DiscountID() [16]byte        // discount domain
	UserType() protocol.UserType //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type DiscountUserType_StorageServices interface {
	Save(dut DiscountUserType) (numbers uint64, err protocol.Error)

	Count(discountID [16]byte) (numbers uint64, err protocol.Error)
	Get(discountID [16]byte, versionOffset uint64) (dut DiscountUserType, numbers uint64, err protocol.Error)

	FindByUserType(userType protocol.UserType, offset, limit uint64) (discountIDs [][16]byte, numbers uint64, err protocol.Error)
}
