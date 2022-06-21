/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherCategory indicate the domain record data fields.
// Use to restrict the voucher on specific category.
type VoucherCategory interface {
	VoucherID() [16]byte  // voucher domain
	Each() uint8          // Each time use
	CategoryID() [16]byte // category domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type VoucherCategory_StorageServices interface {
	Save(vc VoucherCategory) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vc VoucherCategory, numbers uint64, err protocol.Error)
}
