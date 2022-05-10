/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type DiscountUser interface {
	DiscountID() [16]byte        // discount domain
	UserID() [16]byte            // user-status domain
	UserType() protocol.UserType //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}
