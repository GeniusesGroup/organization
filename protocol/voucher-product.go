/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherProduct restrict use voucher on specific product or belong to specific category.
type VoucherProduct interface {
	VoucherID() [16]byte  // voucher domain
	Each() uint8          // Each time use
	QuiddityID() [16]byte // quiddity domain. specific product or specific category.
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
