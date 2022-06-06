/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// QuiddityIcon indicate the domain record data fields.
// Icon or emoji use in many domains e.g. category, departments, ...
type QuiddityIcon interface {
	QuiddityID() [16]byte // quiddity domain
	Icon() [16]byte       // object domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}
