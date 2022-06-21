/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherUserType indicate the domain record data fields.
// Use to restrict the voucher usage by specific user type.
type VoucherUserType interface {
	VoucherID() [16]byte         // voucher domain
	UserType() protocol.UserType //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type VoucherUserType_StorageServices interface {
	Save(vut VoucherUserType) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vut VoucherUserType, numbers uint64, err protocol.Error)

	FilterByUserType(userType protocol.UserType, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}
