/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type VoucherPrice interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	Price() int64        // Max discount price per use
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
