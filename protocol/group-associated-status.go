/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// GroupAssociatedStatus indicate the domain record data fields
type GroupAssociatedStatus interface {
	GroupID() [16]byte              // group-status domain
	UserID() [16]byte               // user-status domain
	Status() GroupAssociated_Status //
	Time() protocol.Time            // Save time
	RequestID() [16]byte            // user-request domain
}

type GroupAssociated_Status uint8

const (
	GroupAssociated_Status_Unset GroupAssociated_Status = iota
	GroupAssociated_Status_Joined
	GroupAssociated_Status_Left
	GroupAssociated_Status_Blocked
	GroupAssociated_Status_Muted
)
