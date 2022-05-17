/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// GroupAssociatedRole indicate the domain record data fields.
type GroupAssociatedRole interface {
	GroupID() [16]byte     // group-status domain
	UserID() [16]byte      // user-status domain
	GroupRoleID() [16]byte // group-role domain
	Time() protocol.Time   // Save time
	RequestID() [16]byte   // user-request domain
}
