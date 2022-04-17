/* For license and copyright information please see LEGAL file in repository */

package org

// CommentMentioned indicate the comment-mentioned domain record data fields.
type CommentMentioned interface {
	UserID() [16]byte    // user-status domain
	CommentID() [16]byte // comment domain
}
