/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentStatus indicate the domain record data fields.
type CommentStatus interface {
	CommentID() [16]byte    // comment domain
	Status() Comment_Status //
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type CommentStatus_StorageServices interface {
	Save(cs CommentStatus) protocol.Error

	Count(commentID [16]byte) (numbers uint64, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (cs CommentStatus, err protocol.Error)
	Last(commentID [16]byte) (cs CommentStatus, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type Comment_Status uint8

const (
	Comment_Status_Unset   Comment_Status = 0
	Comment_Status_Created Comment_Status = (1 << iota)
	Comment_Status_Edited
	Comment_Status_Locked
	Comment_Status_Hidden // use also for unpin a comment
	Comment_Status_Deleted
	Comment_Status_Blocked
)
