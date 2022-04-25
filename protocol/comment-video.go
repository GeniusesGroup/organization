/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentVideo indicate the comment-video domain record data fields.
type CommentVideo interface {
	CommentID() [16]byte // comment domain
	ObjectID() [16]byte  // object domain
	Text() string        // Text With Style. HTML & CSS is more expressive than markdown, suggest use them in article text to style text.
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
