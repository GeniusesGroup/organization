/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentPin indicate the domain record data fields.
type CommentPin interface {
	CommentID() [16]byte // comment domain
	GroupID() [16]byte   // group domain
	PinedID() [16]byte   // comment domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type CommentPin_StorageServices interface {
	Save(cp CommentPin) protocol.Error

	Get(commentID [16]byte) (cp CommentPin, err protocol.Error)

	FindByGroupID(groupID [16]byte, offset, limit uint64) (commentIDs [][16]byte, numbers uint64, err protocol.Error)
}
