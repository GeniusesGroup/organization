/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentObject indicate the comment-object domain record data fields.
// Object can be any such as image, video, voice, ...
type CommentObject interface {
	CommentID() [16]byte // comment domain
	ObjectID() [16]byte  // object domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
