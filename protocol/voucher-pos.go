/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherPOS indicate the domain record data fields.
type VoucherPOS interface {
	VoucherID() [16]byte // voucher domain
	PosID() [16]byte     // pos domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherPOS_StorageServices interface {
	Save(vp VoucherPOS) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vp VoucherPOS, numbers uint64, err protocol.Error)

	FindByPOS(posID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}
