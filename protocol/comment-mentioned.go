/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentMentioned indicate the comment-mentioned domain record data fields.
type CommentMentioned interface {
	UserID() [16]byte    // user-status domain
	CommentID() [16]byte // comment domain
	Time() protocol.Time // Save time
}
