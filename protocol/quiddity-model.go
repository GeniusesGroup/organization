/* For license and copyright information please see LEGAL file in repository */

package org

import "../libgo/protocol"

// QuiddityModel is the 3D model of the quiddity
type QuiddityModel interface {
	QuiddityID() [16]byte // quiddity domain
	ObjectID() [32]byte   // object domain. ThingModel
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
