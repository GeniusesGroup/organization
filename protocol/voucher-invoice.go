/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherInvoice indicate the domain record data fields.
type VoucherInvoice interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	MinPrice() int64     // Minimum invoice price
	MinAmount() uint64   // Minimum invoice product numbers
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherInvoice_StorageServices interface {
	Save(vi VoucherInvoice) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vi VoucherInvoice, numbers uint64, err protocol.Error)
}
