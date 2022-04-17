/* For license and copyright information please see LEGAL file in repository */

package org

// CommentVideo indicate the comment-video domain record data fields.
type CommentVideo interface {
	CommentID() [16]byte // comment domain
	ObjectID() [16]byte
	Text() string        // Text With Style. HTML & CSS is more expressive than markdown, suggest use them in article text to style text.
	RequestID() [16]byte // user-request domain
}
