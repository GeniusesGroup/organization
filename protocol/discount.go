/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"../libgo/protocol"
)

// Discount to reduce invoices price automatically by some rules know also as sales, ...
type Discount interface {
	ID() [16]byte        // product-auction domain
	OrgID() [16]byte     // user-status domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
