/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type ProductReserveValidity interface {
	ProductID() [16]byte         // product-status domain
	Duration() protocol.Duration // from add to invoice time
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}
