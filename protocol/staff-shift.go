/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// StaffShift indicate the domain record data fields.
type StaffShift interface {
	StaffID() [16]byte    // staff-status domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
