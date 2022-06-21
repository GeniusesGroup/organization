/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherGroup indicate the domain record data fields.
type VoucherGroup interface {
	VoucherID() [16]byte // voucher domain
	GroupID() [16]byte   // group domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherGroup_StorageServices interface {
	Save(vg VoucherGroup) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vg VoucherGroup, numbers uint64, err protocol.Error)

	FindByGroup(groupID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}
