/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type ContentStatus interface {
	ContentID() [16]byte          // content domain
	Status() ContentStatus_Status //
	Time() protocol.Time          // Save time
	RequestID() [16]byte          // user-request domain
}

type ContentStatus_StorageServices interface {
	Save(cs ContentStatus) protocol.Error

	Count(contentID [16]byte) (numbers uint64, err protocol.Error)
	Get(contentID [16]byte, versionOffset uint64) (cs ContentStatus, err protocol.Error)
	Last(contentID [16]byte) (cs ContentStatus, numbers uint64, err protocol.Error)
}

type ContentStatus_Status uint16

const (
	ContentStatus_Status_Unset ContentStatus_Status = iota
	ContentStatus_Status_Created
	ContentStatus_Status_Locked
	ContentStatus_Status_Hidden
	ContentStatus_Status_Deleted
	ContentStatus_Status_Blocked
)
