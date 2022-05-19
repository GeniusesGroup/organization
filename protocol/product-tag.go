/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type ProductTag interface {
	Tag() string         //
	ProductID() [16]byte // product-status domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
