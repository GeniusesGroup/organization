/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// GroupStatus indicate the domain record data fields.
type GroupStatus interface {
	GroupID() [16]byte    // group-status domain
	Status() Group_Status //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type GroupStatus_StorageServices interface {
	Save(ur GroupStatus) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (ur GroupStatus, err protocol.Error)
	Last(groupID [16]byte) (ur GroupStatus, err protocol.Error)
}

type Group_Status uint

const (
	Group_Status_Unset Group_Status = iota
	Group_Status_Registered
	Group_Status_Removed
	Group_Status_Blocked
)
