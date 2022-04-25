/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// VoucherUser restrict voucher usage by specific user or user type.
type VoucherUser interface {
	VoucherID() [16]byte         // voucher domain
	UserID() [16]byte            // user-status domain.
	UserType() protocol.UserType //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}
