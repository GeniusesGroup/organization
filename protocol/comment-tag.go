/* For license and copyright information please see LEGAL file in repository */

package org

// CommentTag or hashtag indicate the comment-tag domain record data fields.
type CommentTag interface {
	Tag() string
	CommentID() [16]byte // comment domain
}
