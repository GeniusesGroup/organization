/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentMentioned indicate the domain record data fields.
type CommentMentioned interface {
	UserID() [16]byte    // user-status domain
	CommentID() [16]byte // comment domain
	Time() protocol.Time // Save time
}

type CommentMentioned_StorageServices interface {
	Save(cm CommentMentioned) protocol.Error

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (cm CommentMentioned, err protocol.Error)
	Last(userID [16]byte) (cm CommentMentioned, numbers uint64, err protocol.Error)
}
