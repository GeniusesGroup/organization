/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentForward indicate the domain record data fields.
type CommentForward interface {
	CommentID() [16]byte   // comment domain
	ForwardedID() [16]byte // comment domain. or ForwardedCommentID
	Time() protocol.Time   // Save time
	RequestID() [16]byte   // user-request domain
}

type CommentForward_StorageServices interface {
	Save(cf CommentForward) protocol.Error

	Get(commentID [16]byte) (cf CommentForward, err protocol.Error)

	FindByForwardedID(forwardedID [16]byte, offset, limit uint64) (commentIDs [][16]byte, numbers uint64, err protocol.Error)
}
