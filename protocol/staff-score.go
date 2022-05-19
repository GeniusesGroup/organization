/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// StaffScore or staff grade
type StaffScore interface {
	StaffID() [16]byte   // staff-status domain
	Score() uint64       //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
