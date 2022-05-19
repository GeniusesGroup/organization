/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// ProductProvider or product operator is to supply a particular service.
type ProductProvider interface {
	ProductID() [16]byte // product-status domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

// related UserIDs by RoleID use to get Coordinate(domain)
