/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type VoucherInvoice interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	MinPrice() int64     // Minimum invoice price
	MinAmount() uint64   // Minimum invoice product numbers
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
