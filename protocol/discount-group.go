/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
)

type DiscountGroup interface {
	DiscountID() [16]byte // discount domain
	GroupID() [16]byte    // group domain
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
