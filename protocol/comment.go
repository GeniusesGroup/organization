/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Comment indicate the comment domain record data fields.
// each comment is immutable record and so use version mechanism to chain comments in a group.
type Comment interface {
	GroupID() [16]byte // comment-group domain.
	CommentID() [16]byte
	CommentType() CommentType
	ReplyTo() [16]byte   // other CommentID. Comment can be reply to other comment in the same group.
	UserID() [16]byte    // user-status domain
	RequestID() [16]byte // user-request domain
}

type CommentType uint8

const (
	CommentType_Unset CommentType = iota
	CommentType_Text
	CommentType_Voice
	CommentType_Call
	CommentType_File

	CommentType_Image
	CommentType_Video
	CommentType_Album

	CommentType_Thread
	CommentType_Forward // Retweet, Share, ...
)

type CommentStorageServices interface {
	Register(c Comment) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (c Comment, err protocol.Error)
	Last(groupID [16]byte) (c Comment, err protocol.Error)

	FindByReplyTo(groupID [16]byte, commentID [16]byte, offset, limit uint64) (versionOffsets []uint64, err protocol.Error)
	FindByUserID(groupID [16]byte, userID [16]byte, offset, limit uint64) (versionOffsets []uint64, err protocol.Error)

	protocol.EventTarget
}
