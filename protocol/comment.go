/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Comment indicate the comment domain record data fields.
// each comment is immutable record and so use version mechanism to chain comments in a group.
type Comment interface {
	GroupID() [16]byte      // comment-group domain.
	CommentID() [16]byte    //
	ReplyTo() [16]byte      // other CommentID. Comment can be reply to other comment in the same group.
	UserID() [16]byte       // user-status domain
	Type() Comment_Type     //
	Status() Comment_Status //
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type Comment_StorageServices interface {
	Save(c Comment) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (c Comment, err protocol.Error)
	Last(groupID [16]byte) (c Comment, err protocol.Error)

	FindByReplyTo(groupID [16]byte, commentID [16]byte, offset, limit uint64) (versionOffsets []uint64, numbers uint64, err protocol.Error)
	FindByUserID(groupID [16]byte, userID [16]byte, offset, limit uint64) (versionOffsets []uint64, numbers uint64, err protocol.Error)

	ListUserGroups(userID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type Comment_Type uint8

const (
	Comment_Type_Unset Comment_Type = iota
	Comment_Type_Text
	Comment_Type_Voice
	Comment_Type_Call
	Comment_Type_File

	Comment_Type_Image
	Comment_Type_Video
	Comment_Type_Album

	Comment_Type_Thread
	Comment_Type_Forward // Retweet, Share, ...

	Comment_Type_NewConnection
)

type Comment_Status uint8

const (
	Comment_Status_Unset  Comment_Status = 0
	Comment_Status_Locked Comment_Status = (1 << iota)
	Comment_Status_Hidden
	Comment_Status_Deleted
	Comment_Status_Blocked
)
