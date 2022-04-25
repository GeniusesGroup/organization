/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentTag or hashtag indicate the comment-tag domain record data fields.
type CommentTag interface {
	Tag() string         //
	CommentID() [16]byte // comment domain
	Time() protocol.Time // Save time
}
