/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentPin indicate the comment-pin domain record data fields.
type CommentPin interface {
	GroupID() [16]byte   // comment-group domain.
	CommentID() [16]byte // comment domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
