/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type DiscountLocation interface {
	DiscountID() [16]byte // discount domain
	LocationID() [16]byte // location domain
	ZoneID() [16]byte     // location domain
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
