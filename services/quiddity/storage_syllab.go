/* For license and copyright information please see LEGAL file in repository */

package quiddity

import (
	"../../libgo/protocol"
)

type quiddityAsSyllab []byte

func (q *quiddityAsSyllab) QuiddityID() [16]byte {}
func (q *quiddityAsSyllab) DomainID() uint64     {}
func (q *quiddityAsSyllab) Time() protocol.Time  {}
func (q *quiddityAsSyllab) RequestID() [16]byte  {}
