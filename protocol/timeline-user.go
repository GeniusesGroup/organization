/* For license and copyright information please see LEGAL file in repository */

package org

// UserTimeline indicate the timeline-user domain record data fields.
type UserTimeline interface {
	UserID() [16]byte         // user-status domain
	CommentGroupID() [16]byte // generate for comment-group & comment domain
	Status() UserTimelineStatus
	RequestID() [16]byte // user-request domain
}

type UserTimelineStatus uint8

const (
	UserTimelineStatus_Unset UserTimelineStatus = iota
	UserTimelineStatus_Locked
	UserTimelineStatus_Hidden
	UserTimelineStatus_Deleted // soft delete
)
