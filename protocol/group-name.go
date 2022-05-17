/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// GroupName indicate the domain record data fields.
// It is not the group title as display name, It is the unique group name.
type GroupName interface {
	GroupID() [16]byte        // group-status domain
	Name() string             // It is not replace of group ID. It usually use to find the group in more human friendly manner.
	Status() GroupName_Status //
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

type GroupName_StorageServices interface {
	Save(gn GroupName) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gn GroupName, err protocol.Error)
	Last(groupID [16]byte) (gn GroupName, err protocol.Error)

	FindByGroupName(groupname string) (groupID [16]byte, err protocol.Error)
}

// GroupName_Status indicate GroupName record status
type GroupName_Status uint8

// GroupName status
const (
	GroupName_Status_Unset    GroupName_Status = 0
	GroupName_Status_Register GroupName_Status = (1 << iota)
	GroupName_Status_Remove
	GroupName_Status_Blocked
	GroupName_Status_Reserved
)
