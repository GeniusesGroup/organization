/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Department indicate the domain record data fields.
type Department interface {
	ID() [16]byte        // quiddity domain. use to get and show locale title.
	ParentID() [16]byte  // DepartmentID, Exist if it isn't top of the organization.
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
