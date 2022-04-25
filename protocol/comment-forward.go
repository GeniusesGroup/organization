/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentForward indicate the comment-forward domain record data fields.
type CommentForward interface {
	CommentID() [16]byte // comment domain
	Forwarded() [16]byte // comment domain. or ForwardedCommentID
	Text() string        // Text With Style. HTML & CSS (No JS) is more expressive than markdown, suggest use them in article text to style text.
	Preview() bool       // Render web links as small widget or not
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
