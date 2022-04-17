/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// CommentThread indicate the comment-thread domain record data fields.
type CommentThread interface {
	CommentID() [16]byte // comment domain
	Name() string
	ArchiveAfter() protocol.Duration
	RequestID() [16]byte // user-request domain
}
