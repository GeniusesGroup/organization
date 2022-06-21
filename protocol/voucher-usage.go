/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherUsage indicate the domain record data fields.
type VoucherUsage interface {
	VoucherID() [16]byte    // product domain
	MaximumInvoice() uint64 //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type VoucherUsage_StorageServices interface {
	Save(vu VoucherUsage) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vu VoucherUsage, numbers uint64, err protocol.Error)
}
