/* For license and copyright information please see LEGAL file in repository */

package org

// TimelinePin
type TimelinePin interface {
	UserID() [16]byte         // user-status domain
	CommentGroupID() [16]byte // comment-group domain
	RequestID() [16]byte      // user-request domain
}
