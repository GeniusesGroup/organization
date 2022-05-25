/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Comment indicate the comment domain record data fields.
// each comment is immutable record and so use version mechanism to chain comments in a group.
type Comment interface {
	GroupID() [16]byte          // group domain
	CommentID() [16]byte        //
	ReplyTo() [16]byte          // other CommentID. Comment can be reply to other comment in the same group.
	UserID() [16]byte           // user-status domain
	Type() Comment_Type         //
	Settings() Comment_Settings //
	Time() protocol.Time        // Save time
	RequestID() [16]byte        // user-request domain
}

type Comment_StorageServices interface {
	Save(c Comment) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (c Comment, err protocol.Error)
	Last(groupID [16]byte) (c Comment, numbers uint64, err protocol.Error)

	FindByReplyTo(groupID [16]byte, commentID [16]byte, offset, limit uint64) (versionOffsets []uint64, numbers uint64, err protocol.Error)
	FindByUserID(groupID [16]byte, userID [16]byte, offset, limit uint64) (versionOffsets []uint64, numbers uint64, err protocol.Error)

	ListUserGroups(userID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

type Comment_Type uint32

const (
	Comment_Type_Unset Comment_Type = 0
	Comment_Type_Text  Comment_Type = (1 << iota)
	Comment_Type_Voice
	Comment_Type_Call
	Comment_Type_File

	Comment_Type_Image
	Comment_Type_Video
	Comment_Type_Album

	Comment_Type_Pin
	Comment_Type_Welcome
	Comment_Type_PhotoUpdated
	Comment_Type_NewConnection // made-friend, new-follower,

	Comment_Type_Map // location
	Comment_Type_Product

	Comment_Type_Thread  // Thread name will save in comment-text
	Comment_Type_Forward // Retweet, Share, ...
)

type Comment_Settings uint32

const (
	Comment_Settings_Unset        Comment_Settings = 0
	Comment_Settings_PreviewLinks Comment_Settings = (1 << iota) // Render web links as small widget or not
	Comment_Settings_Forwardable                                 // Allow to forward it
)
